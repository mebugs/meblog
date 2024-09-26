package platService

import (
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"siteol.com/smart/src/service/auth/authService"
)

// getPermissionRouters 获取权限路由集 withRouter 是否提取路由信息
func getPermissionRouters(traceID string, permissionId uint64, withRouter bool) (routerIds []uint64, routers []*platModel.RouterPageRes, err error) {
	// 补充查询关联的路由和路由ID
	routerIds, err = platDB.PermissionRouter{}.GetRouterIds(permissionId)
	if err != nil {
		log.WarnTF(traceID, "GetRouterIdsByPermissionId By %d Fail . Err Is : %v", permissionId, err)
		return
	}
	list := make([]*platDB.Router, 0)
	// 关联查询路由字段
	if withRouter && len(routerIds) > 0 {
		listR, errR := platDB.RouterTable.GetByIds(routerIds)
		if errR != nil {
			err = errR
			log.WarnTF(traceID, "GetRouterByRouterIds By %v Fail . Err Is : %v", routerIds, err)
			return
		}
		list = listR
	}
	routers = platModel.ToRouterPageRes(list)
	return
}

// syncPermissionRouters 编辑权限对应的路由
func syncPermissionRouters(traceID string, permissionId uint64, routerIds []uint64, editFlag bool) (err error) {
	if editFlag {
		// 移除当前权限的路由
		err = platDB.PermissionRouterTable.Executor().DeleteByPermissionId(permissionId)
		if err != nil {
			log.ErrorTF(traceID, "DeleteByPermissionId By %d Fail . Err Is : %v", permissionId, err)
			return
		}
	}
	// 重新插入路由关系
	if len(routerIds) > 0 {
		permissionRouters := make([]platDB.PermissionRouter, len(routerIds))
		for i, item := range routerIds {
			permissionRouters[i] = platDB.PermissionRouter{
				PermissionId: permissionId,
				RouterId:     item,
			}
		}
		err = platDB.PermissionRouterTable.InsertBatch(&permissionRouters)
		if err != nil {
			log.ErrorTF(traceID, "InsertBatchPermissionRouter By PermissionId %d Fail . Err Is : %v", permissionId, err)
			return
		}
	}
	if editFlag {
		go func() {
			// 获取权限绑定的角色
			roleIds, err := platDB.RolePermissionTable.Executor().GetRoleIds(permissionId)
			if err != nil {
				log.WarnTF(traceID, "RefreshAuthCache By permissionId %d Fail . Err Is : %v", permissionId, err)
			}
			// 获取角色关联的账号
			if len(roleIds) == 0 {
				return
			}
			// 如果角色被选择，则反向通知账号需要权限刷新
			accountIds, err := platDB.AccountRoleTable.Executor().GetAccountIdsWithRoleIds(roleIds)
			if err != nil {
				log.WarnTF(traceID, "RefreshAuthCache By RoleIds %v Fail . Err Is : %v", roleIds, err)
			}
			// 通知账号权限有刷新
			authService.RefreshAuthCacheByAccounts(traceID, accountIds)
		}()
	}
	return
}
