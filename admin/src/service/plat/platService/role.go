package platService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/cacheModel"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/common/mysql/platDB"
)

// AddRole 创建角色
func AddRole(traceID string, req *platModel.RoleAddReq) *baseModel.ResBody {
	// 创建对象初始化
	dbReq := req.ToDbReq()
	err := platDB.RoleTable.InsertOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddRole Fail . Err Is : %v", err)
		// 解析数据库错误
		return checkRoleDBErr(err)
	}
	// 处理权限选择
	err = syncRolePermissions(traceID, dbReq.Id, req.PermissionIds, req.HalfPermissionIds, false)
	if err != nil {
		// 移除当前值
		errD := platDB.RoleTable.DeleteOne(dbReq.Id)
		if errD != nil {
			log.ErrorTF(traceID, "AddRole Rollback Fail . Err Is : %v", err)
		}
		return baseModel.Fail(constant.RolePermissionNG)
	}
	// 同步角色缓存
	go func() { _ = cacheModel.SyncRoleCache(traceID) }()
	return baseModel.Success(constant.RoleAddSS, true)
}

// PageRole 查询角色分页
func PageRole(traceID string, req *platModel.RolePageReq) *baseModel.ResBody {
	// 查询分页
	total, list, err := platDB.RoleTable.Page(rolePageQuery(req))
	if err != nil {
		log.ErrorTF(traceID, "PageRole Fail . Err Is : %v", err)
		return baseModel.Fail(constant.RoleGetNG)
	}
	return baseModel.SuccessUnPop(baseModel.SetPageRes(platModel.ToRolePageRes(list), total))
}

// GetRole 角色详情
func GetRole(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := platDB.RoleTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetRole Fail . Err Is : %v", err)
		return baseModel.Fail(constant.RoleGetNG)
	}
	rolePermissions, err := getRolePermissions(traceID, req.Id)
	if err != nil {
		return baseModel.Fail(constant.RoleGetNG)
	}
	return baseModel.SuccessUnPop(platModel.ToRoleGetRes(&res, rolePermissions))
}

// EditRole 编辑角色
func EditRole(traceID string, req *platModel.RoleEditReq) *baseModel.ResBody {
	dbReq, err := platDB.RoleTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetRole Fail . Err Is : %v", err)
		return baseModel.Fail(constant.RoleGetNG)
	}
	// 先处理权限选择
	err = syncRolePermissions(traceID, dbReq.Id, req.PermissionIds, req.HalfPermissionIds, true)
	if err != nil {
		return baseModel.Fail(constant.RolePermissionNG)
	}
	// 对象更新
	req.ToDbReq(&dbReq)
	err = platDB.RoleTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "EditRole %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return checkRoleDBErr(err)
	}
	// 同步角色缓存
	go func() { _ = cacheModel.SyncRoleCache(traceID) }()
	return baseModel.Success(constant.RoleEditSS, true)
}

// DelRole 角色封存
func DelRole(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	dbReq, err := platDB.RoleTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetRole Fail . Err Is : %v", err)
		return baseModel.Fail(constant.RoleGetNG)
	}
	// 角色禁止刪除
	if dbReq.Mark == constant.StatusLock {
		log.ErrorTF(traceID, "DelRole %d Fail . Can not Edit", dbReq.Id)
		return baseModel.Fail(constant.RoleMarkNG)
	}
	// 先删除权限
	err = syncRolePermissions(traceID, dbReq.Id, nil, nil, true)
	if err != nil {
		return baseModel.Fail(constant.RolePermissionNG)
	}
	// 删除角色关联账号
	err = platDB.AccountRole{}.DeleteByRoleId(dbReq.Id)
	if err != nil {
		log.ErrorTF(traceID, "DeleteAccountRoleByRole %d Fail . Err Is : %v", dbReq.Id, err)
		return baseModel.Fail(constant.RoleDelNG)
	}
	// 删除角色
	err = platDB.RoleTable.DeleteOne(dbReq.Id)
	if err != nil {
		log.ErrorTF(traceID, "DelRole %d Fail . Err Is : %v", dbReq.Id, err)
		return baseModel.Fail(constant.RoleDelNG)
	}
	// 同步角色缓存
	go func() { _ = cacheModel.SyncRoleCache(traceID) }()
	return baseModel.Success(constant.RoleDelSS, true)
}

// ListRole 角色列表
func ListRole(traceID string) *baseModel.ResBody {
	roleMap, err := cacheModel.GetRoleCache(traceID)
	if err != nil {
		return baseModel.Fail(constant.RoleGetNG)
	}
	roleList := make([]baseModel.SelectNumRes, 0)
	for k, v := range roleMap {
		roleList = append(roleList, baseModel.SelectNumRes{
			Label: v,
			Value: k,
		})
	}
	return baseModel.SuccessUnPop(roleList)
}
