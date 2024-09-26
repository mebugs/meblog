package platService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/cacheModel"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"sort"
)

// AddPermission 创建权限
func AddPermission(traceID string, req *platModel.PermissionAddReq) *baseModel.ResBody {
	// 创建对象初始化
	dbReq := req.ToDbReq()
	err := platDB.PermissionTable.InsertOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddPermission Fail . Err Is : %v", err)
		// 解析数据库错误
		return checkPermissionDBErr(err)
	}
	// 新增权限路由
	err = syncPermissionRouters(traceID, dbReq.Id, req.RouterIds, false)
	if err != nil {
		// 移除当前值
		errD := platDB.PermissionTable.DeleteOne(dbReq.Id)
		if errD != nil {
			log.ErrorTF(traceID, "AddPermission Rollback Fail . Err Is : %v", err)
		}
		return baseModel.Fail(constant.PermissionRouterNG)
	}
	go func() { _ = cacheModel.SyncPermissionCache(traceID) }()
	return baseModel.Success(constant.PermissionAddSS, true)
}

// TreePermission 查询权限树
func TreePermission(traceID string) *baseModel.ResBody {
	// 查询根节点
	rootPerm, err := platDB.PermissionTable.GetOneById(1)
	if err != nil {
		log.ErrorTF(traceID, "TreePermission GetRoot Fail . Err Is : %s", err)
		return baseModel.Fail(constant.PermissionGetNG)
	}
	// 创建树节点
	treeNode := &baseModel.Tree{
		Title:    rootPerm.Name,
		Key:      rootPerm.Alias,
		Children: nil,
		Expand:   rootPerm.Static == constant.StatusOpen, // 拓展用于选择框置灰
		Level:    rootPerm.Level,
		Id:       rootPerm.Id,
	}
	// 递归权限树
	_ = recursionPermissionTree(traceID, treeNode)
	trees := []*baseModel.Tree{treeNode}
	return baseModel.SuccessUnPop(trees)
}

// GetPermission 权限详情
func GetPermission(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := platDB.PermissionTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetPermission Fail . Err Is : %v", err)
		return baseModel.Fail(constant.PermissionGetNG)
	}
	// 获取路由集数据
	routerIds, routers, _ := getPermissionRouters(traceID, res.Id, true)
	return baseModel.SuccessUnPop(platModel.ToPermissionGetRes(&res, routerIds, routers))
}

// EditPermission 编辑权限
func EditPermission(traceID string, req *platModel.PermissionEditReq) *baseModel.ResBody {
	dbReq, err := platDB.PermissionTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetPermission Fail . Err Is : %v", err)
		return baseModel.Fail(constant.PermissionGetNG)
	}
	// 先处理路由数据
	err = syncPermissionRouters(traceID, req.Id, req.RouterIds, true)
	if err != nil {
		return baseModel.Fail(constant.PermissionRouterNG)
	}
	// 对象更新
	req.ToDbReq(&dbReq)
	err = platDB.PermissionTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "EditPermission %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return checkPermissionDBErr(err)
	}
	go func() { _ = cacheModel.SyncPermissionCache(traceID) }()
	return baseModel.Success(constant.PermissionEditSS, true)
}

// DelPermission 权限封存
func DelPermission(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	dbReq, err := platDB.PermissionTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetPermissionFail . Err Is : %v", err)
		return baseModel.Fail(constant.PermissionGetNG)
	}
	// 权限禁止刪除
	if dbReq.Mark == constant.StatusLock {
		log.ErrorTF(traceID, "DelPermission %d Fail . Can not Edit", dbReq.Id)
		return baseModel.Fail(constant.PermissionMarkNG)
	}
	// 先删除路由关联，传入空白数组只删不加
	err = syncPermissionRouters(traceID, req.Id, nil, true)
	if err != nil {
		return baseModel.Fail(constant.PermissionRouterNG)
	}
	// 删除权限关联角色
	err = platDB.RolePermissionTable.Executor().DeleteByPermissionId(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "DelRolePermissionByPermission %d Fail . Err Is : %v", dbReq.Id, err)
		return baseModel.Fail(constant.PermissionDelNG)
	}
	// 删除权限
	err = platDB.PermissionTable.DeleteOne(dbReq.Id)
	if err != nil {
		log.ErrorTF(traceID, "DelPermission %d Fail . Err Is : %v", dbReq.Id, err)
		return baseModel.Fail(constant.PermissionDelNG)
	}
	go func() { _ = cacheModel.SyncPermissionCache(traceID) }()
	return baseModel.Success(constant.PermissionDelSS, true)
}

// BroPermission 同级权限列表
func BroPermission(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	dbReq, err := platDB.PermissionTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetPermissionFail . Err Is : %v", err)
		return baseModel.Fail(constant.PermissionGetNG)
	}
	var bros platDB.PermissionArray
	bros, err = platDB.PermissionTable.GetByObject(&platDB.Permission{Pid: dbReq.Pid})
	if err != nil {
		log.ErrorTF(traceID, "GetBroPermissionFail . Err Is : %v", err)
		return baseModel.Fail(constant.PermissionGetNG)
	}
	sort.Sort(bros)
	return baseModel.SuccessUnPop(platModel.ToPermissionBroRes(bros))
}

// SortPermission 权限排序
func SortPermission(traceID string, req *[]*baseModel.SortReq) *baseModel.ResBody {
	reqObj := *req
	if len(reqObj) == 0 {
		return baseModel.SysErr
	}
	err := platDB.PermissionTable.SortWithTransaction(reqObj)
	if err != nil {
		log.ErrorTF(traceID, "SortPermission Fail . Err Is : %v", err)
		// 解析数据库错误
		return checkPermissionDBErr(err)
	}
	return baseModel.Success(constant.PermissionSortSS, true)
}
