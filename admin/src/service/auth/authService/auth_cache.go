package authService

import (
	"fmt"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/cacheModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"siteol.com/smart/src/common/redis"
	"time"
)

// RefreshAuthCacheByAccounts 刷新受影响的账号权限
func RefreshAuthCacheByAccounts(traceID string, accountIds []uint64) {
	if len(accountIds) == 0 {
		return
	}
	log.InfoTF(traceID, "RefreshAuthCacheByAccounts By %vStart", accountIds)
	accounts, err := platDB.AccountTable.GetByIds(accountIds)
	if err != nil {
		log.WarnTF(traceID, "RefreshAuthCacheByAccounts By %v Fail . Err Is : %v", accountIds, err)
		return
	}
	// 获取账号的登陆信息
	for _, account := range accounts {
		// 检索出历史存活的登陆数据
		records, err := platDB.LoginRecordTable.Executor().GetOutRangeRecords(account.Id, 0)
		if err != nil {
			log.WarnTF(traceID, "RefreshAuthCacheByAccounts GetLoginRecords Fail . Err Is : %v", err)
			continue
		}
		outIds := make([]uint64, 0)
		// 遍历可能得在线记录
		for _, record := range records {
			// 尝试获取数据
			outTime, err := redis.GetTTL(fmt.Sprintf(constant.CacheAuth, record.Token))
			// 账号状态被关闭
			if err != nil || account.Status != constant.StatusOpen {
				// 踢出登陆
				outIds = append(outIds, record.Id)
				continue
			}
			// 以剩余时间更新授权缓存
			_, err = makeAuthCacheByOutTime(traceID, record.Token, *account, nil, &outTime)
			if err != nil {
				log.WarnTF(traceID, "RefreshAuthCacheByAccounts MakeAuthCache Fail . Err Is : %v", err)
			}
		}
		// 如果存在需要下线的Token
		now := time.Now()
		// 批量更新
		err = platDB.LoginRecordTable.UpdateByIds(outIds, map[string]any{
			"mark":      constant.StatusClose, // 被动登出
			"update_at": &now,
		})
		if err != nil {
			log.WarnTF(traceID, "RefreshAuthCacheByAccounts UpdateLoginRecords Fail . Err Is : %v", err)
		}
	}
	log.InfoTF(traceID, "RefreshAuthCacheByAccounts By %v Done", accountIds)
}

// makeAuthCache 生成授权缓存
func makeAuthCache(traceID, token string, account platDB.Account, conf *cacheModel.CacheSysConfig) (cacheAuth *cacheModel.CacheAuth, err error) {
	return makeAuthCacheByOutTime(traceID, token, account, conf, nil)
}

// makeAuthCacheByOutTime 生成授权缓存
func makeAuthCacheByOutTime(traceID, token string, account platDB.Account, conf *cacheModel.CacheSysConfig, outTime *time.Duration) (cacheAuth *cacheModel.CacheAuth, err error) {
	needReset := false
	if account.PwdExpTime != nil {
		needReset = time.Now().After(*account.PwdExpTime)
	}
	// 基础权限数据
	cacheAuth = &cacheModel.CacheAuth{
		AccountId:    account.Id,
		Name:         account.Name,
		NeedResetPwd: needReset, // 是否需要重置密码
	}
	// 角色处理
	setCacheRole(traceID, account, cacheAuth)
	// 权限处理
	setCachePermission(traceID, cacheAuth)
	// 路由处理
	setCacheRouter(traceID, cacheAuth)
	// 数据权限处理，如果是部门&子部门需要递归子部门列表
	setCacheDataPermission(traceID, account, cacheAuth)
	// 指定超时时间
	if outTime != nil {
		err = redis.SetByTimeDuration(fmt.Sprintf(constant.CacheAuth, token), cacheAuth, *outTime)
	} else if conf != nil {
		// 根据配置，写入Redis
		outMs := uint64(0)
		if conf.LogoutSwitch {
			outMs = conf.LogoutLimit
		}
		// 写入缓存
		err = redis.Set(fmt.Sprintf(constant.CacheAuth, token), cacheAuth, outMs)
	}
	return
}

// setCacheDataPermission 设置数据权限相关信息
func setCacheDataPermission(traceID string, account platDB.Account, cacheAuth *cacheModel.CacheAuth) {
	// 账号设定为继承部门，可能涉及部门数据权限
	if account.PermissionType == constant.DataPermissionTree {
		dept, err := platDB.DeptTable.GetOneById(account.DeptId)
		if err != nil {
			log.WarnTF(traceID, "GetAccountDept %d Fail . Err Is : %v", account.DeptId, err)
			// 出错则认为是个人权限
			cacheAuth.DataPermissionType = constant.DataPermissionSelf
		} else {
			cacheAuth.DataPermissionType = dept.PermissionType
		}
	}
	// 不同的权限类型分别处理
	// 数据权限分类（如果继承部门，查询部门） 权限类型，枚举：0_继承部门、本级与子集 1_本部门 2_本人 3_全局
	switch cacheAuth.DataPermissionType {
	case constant.DataPermissionTree: // 本级与子集
		cacheAuth.DeptIds = getDeptIds(traceID, account.DeptId)
	case constant.DataPermissionDept, constant.DataPermissionSelf: // 本级/个人
		cacheAuth.DeptIds = []uint64{account.DeptId}
	case constant.DataPermissionAll: // 所有的无需处理
	}
}

// getDeptIds 获取用户具备权限的部门ID
func getDeptIds(traceID string, rootDeptId uint64) (deptIds []uint64) {
	treeNode, err := cacheModel.GetDeptTreeCache(traceID)
	if err != nil {
		deptIds = []uint64{rootDeptId}
		return
	}
	// 提取所属部门的根位置
	rootDept := recursionRootDept(rootDeptId, treeNode)
	if rootDept == nil {
		deptIds = []uint64{rootDeptId}
		return
	}
	// 遍历本部门&子部门
	deptIds = make([]uint64, 0)
	deptIds = recursionDept(deptIds, rootDept)
	return
}

// recursionDept 递归树获取所有ID
func recursionDept(deptIds []uint64, treeNode *baseModel.Tree) []uint64 {
	// 循环子集
	if len(treeNode.Children) > 0 {
		for _, item := range treeNode.Children {
			return recursionDept(deptIds, item)
		}
	} else {
		deptIds = append(deptIds, treeNode.Id)
	}
	return deptIds
}

// recursionRootDept 递归树获取根节点
func recursionRootDept(deptId uint64, treeNode *baseModel.Tree) *baseModel.Tree {
	if deptId == treeNode.Id {
		// 获得根节点
		return treeNode
	}
	// 循环子集
	if len(treeNode.Children) > 0 {
		for _, item := range treeNode.Children {
			getRoot := recursionRootDept(deptId, item)
			if getRoot != nil {
				return getRoot
			}
		}
	}
	return nil
}

// setCacheRouter 设置路由相关信息
func setCacheRouter(traceID string, cacheAuth *cacheModel.CacheAuth) {
	if len(cacheAuth.PermissionIds) == 0 {
		return
	}
	// 基于权限查询路由，白名单路由无需添加到个人，中间件默认开放
	routerIds, errP := platDB.PermissionRouterTable.Executor().GetRouterIdsWithByPermissionIds(cacheAuth.PermissionIds)
	if errP != nil {
		log.WarnTF(traceID, "GetRouterIdsWithByPermissionIds Fail . Err Is : %v", errP)
	}
	if len(routerIds) == 0 {
		return
	}
	routerUrlMap, err := cacheModel.GetRouterUrlsCache(traceID)
	if err != nil {
		return
	}
	routerUrls := make([]string, len(routerIds))
	for i, item := range routerIds {
		routerUrls[i] = routerUrlMap[item]
	}
	// 设置路由
	cacheAuth.Routers = routerUrls
}

// setCachePermission 设置权限相关信息
func setCachePermission(traceID string, cacheAuth *cacheModel.CacheAuth) {
	// 查询权限
	permIds := make([]uint64, 0)
	if len(cacheAuth.RoleIds) > 0 {
		// 基于角色查询权限
		permIdsQ, err := platDB.RolePermissionTable.Executor().GetPermissionIdsWithRoleIds(cacheAuth.RoleIds)
		if err != nil {
			log.WarnTF(traceID, "GetPermissionIdsWithRolIds Fail . Err Is : %v", err)
		} else {
			permIds = permIdsQ
		}
	}
	// 获取默认权限
	normalPermIds, _ := cacheModel.GetPermissionNormalIdsCache(traceID)
	// 取并集
	allPermIds := make([]uint64, 0)
	permMap := make(map[uint64]bool, 0)
	for _, pi := range permIds {
		if _, ok := permMap[pi]; !ok {
			allPermIds = append(allPermIds, pi)
			permMap[pi] = true
		}
	}
	for _, pi := range normalPermIds {
		if _, ok := permMap[pi]; !ok {
			allPermIds = append(allPermIds, pi)
			permMap[pi] = true
		}
	}
	if len(allPermIds) == 0 {
		return
	}
	permAliasMap, err := cacheModel.GetPermissionCache(traceID)
	if err != nil {
		return
	}
	permAlias := make([]string, len(allPermIds))
	for i, item := range allPermIds {
		permAlias[i] = permAliasMap[item]
	}
	// 设置权限和权限别名
	cacheAuth.PermissionIds = allPermIds
	cacheAuth.PermissionKeys = permAlias
}

// setCacheRole 设置角色相关信息
func setCacheRole(traceID string, account platDB.Account, cacheAuth *cacheModel.CacheAuth) {
	// 查询角色
	roleIds, err := platDB.AccountRoleTable.Executor().GetRoleIds(account.Id)
	if err != nil {
		log.WarnTF(traceID, "GetAccountRole %d Fail . Err Is : %v", account.Id, err)
	}
	// 设置角色和角色名
	cacheAuth.RoleIds = roleIds
	if len(roleIds) > 0 {
		roleMap, err := cacheModel.GetRoleCache(traceID)
		if err != nil {
			return
		}
		roleNames := make([]string, len(roleIds))
		for i, item := range roleIds {
			roleNames[i] = roleMap[item]
		}
		cacheAuth.RoleNames = roleNames
	}
}
