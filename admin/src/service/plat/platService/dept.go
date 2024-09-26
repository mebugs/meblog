package platService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/cacheModel"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"sort"
	"time"
)

// AddDept 创建集团部门
func AddDept(traceID string, req *platModel.DeptAddReq) *baseModel.ResBody {
	// 创建对象初始化
	dbReq := req.ToDbReq()
	err := platDB.DeptTable.InsertOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddDept Fail . Err Is : %v", err)
		// 解析数据库错误
		return checkDeptDBErr(err)
	}
	go func() { _ = cacheModel.SyncDeptTreeCache(traceID) }()
	return baseModel.Success(constant.DeptAddSS, true)
}

// TreeDept 查询集团树
func TreeDept(traceID string) *baseModel.ResBody {
	treeNode, err := cacheModel.GetDeptTreeCache(traceID)
	if err != nil {
		return baseModel.Fail(constant.DeptGetNG)
	}
	trees := []*baseModel.Tree{treeNode}
	return baseModel.SuccessUnPop(trees)
}

// GetDept 集团部门详情
func GetDept(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := platDB.DeptTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetDept Fail . Err Is : %v", err)
		return baseModel.Fail(constant.DeptGetNG)
	}
	// 获取部门账号
	accounts, err := getDeptAccounts(res.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetDeptAccount %d Fail . Err Is : %v", req.Id, err)
		return baseModel.Fail(constant.DeptGetNG)
	}
	return baseModel.SuccessUnPop(platModel.ToDeptGetRes(&res, accounts))
}

// EditDept 编辑集团部门
func EditDept(traceID string, req *platModel.DeptEditReq) *baseModel.ResBody {
	dbReq, err := platDB.DeptTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetDept Fail . Err Is : %v", err)
		return baseModel.Fail(constant.DeptGetNG)
	}
	// 对象更新
	req.ToDbReq(&dbReq)
	err = platDB.DeptTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "EditDept %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return checkDeptDBErr(err)
	}
	go func() { _ = cacheModel.SyncDeptTreeCache(traceID) }()
	return baseModel.Success(constant.DeptEditSS, true)
}

// DelDept 集团部门移除
func DelDept(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	dbReq, err := platDB.DeptTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetDept Fail . Err Is : %v", err)
		return baseModel.Fail(constant.DeptGetNG)
	}
	// 集团部门禁止刪除
	if dbReq.Mark == constant.StatusLock {
		log.ErrorTF(traceID, "DelDept %d Fail . Can not Edit", dbReq.Id)
		return baseModel.Fail(constant.DeptMarkNG)
	}
	// 如果部门下存在子部门禁止移除
	childCount, err := platDB.DeptTable.CountByObject(&platDB.Dept{Pid: dbReq.Id})
	if err != nil {
		log.ErrorTF(traceID, "CountChildDept Fail . Err Is : %v", err)
		return baseModel.Fail(constant.DeptGetNG)
	}
	if childCount > 0 {
		log.ErrorTF(traceID, "DelDept Fail . Err Is : Has %d Child", childCount)
		return baseModel.Fail(constant.DeptDelChildNG)
	}
	// 如果部门下存在员工禁止删除
	accCount, err := platDB.DeptTable.CountByObject(&platDB.Account{DeptId: dbReq.Id})
	if err != nil {
		log.ErrorTF(traceID, "CountAccountDept Fail . Err Is : %v", err)
		return baseModel.Fail(constant.DeptGetNG)
	}
	if accCount > 0 {
		log.ErrorTF(traceID, "DelDept Fail . Err Is : Has %d Account", accCount)
		return baseModel.Fail(constant.DeptDelAccountNG)
	}
	err = platDB.DeptTable.DeleteOne(dbReq.Id)
	if err != nil {
		log.ErrorTF(traceID, "DelDept %d Fail . Err Is : %v", dbReq.Id, err)
		return baseModel.Fail(constant.DeptDelNG)
	}
	go func() { _ = cacheModel.SyncDeptTreeCache(traceID) }()
	return baseModel.Success(constant.DeptDelSS, true)
}

// BroDept 同级部门列表
func BroDept(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	dbReq, err := platDB.DeptTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetDeptFail . Err Is : %v", err)
		return baseModel.Fail(constant.DeptGetNG)
	}
	var bros platDB.DeptArray
	bros, err = platDB.DeptTable.GetByObject(&platDB.Dept{Pid: dbReq.Pid})
	if err != nil {
		log.ErrorTF(traceID, "GetBroDeptFail . Err Is : %v", err)
		return baseModel.Fail(constant.DeptGetNG)
	}
	sort.Sort(bros)
	return baseModel.SuccessUnPop(platModel.ToDeptBroRes(bros))
}

// SortDept 部门排序
func SortDept(traceID string, req *[]*baseModel.SortReq) *baseModel.ResBody {
	reqObj := *req
	if len(reqObj) == 0 {
		return baseModel.SysErr
	}
	err := platDB.DeptTable.SortWithTransaction(reqObj)
	if err != nil {
		log.ErrorTF(traceID, "SortDept Fail . Err Is : %v", err)
		// 解析数据库错误
		return checkDeptDBErr(err)
	}
	go func() { _ = cacheModel.SyncDeptTreeCache(traceID) }()
	return baseModel.Success(constant.DeptSortSS, true)
}

// ToDept 部门迁移
func ToDept(traceID string, req *platModel.DeptToReq) *baseModel.ResBody {
	dbReq, err := platDB.DeptTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetDept %d Fail . Err Is : %v", req.Id, err)
		return baseModel.Fail(constant.DeptGetNG)
	}
	_, err = platDB.DeptTable.GetOneById(req.ToId)
	if err != nil {
		log.ErrorTF(traceID, "GetToDept %d Fail . Err Is : %v", req.ToId, err)
		return baseModel.Fail(constant.DeptGetNG)
	}
	now := time.Now()
	// 0_并入（子部门形式），修改本部门的PID即可
	// 1_移交（部门保留，成员和子部门移交给新部门），修改本部门的直系部门的上级部门、成员所属部门
	if req.ToType == constant.StatusOpen {
		dbReq.Pid = req.ToId
		dbReq.UpdateAt = &now
		err = platDB.DeptTable.UpdateOne(dbReq)
		if err != nil {
			log.ErrorTF(traceID, "DeptToDept %d %d Fail . Err Is : %v", req.Id, req.ToId, err)
			return baseModel.Fail(constant.DeptToNG)
		}
	} else {
		// 先迁账号
		err = platDB.AccountTable.Executor().ToNewDept(req.Id, req.ToId)
		if err != nil {
			log.ErrorTF(traceID, "DeptToDept %d %d Account Fail . Err Is : %v", req.Id, req.ToId, err)
			return baseModel.Fail(constant.DeptToNG)
		}
		// 迁子部门
		err = platDB.DeptTable.Executor().ToNewDept(req.Id, req.ToId)
		if err != nil {
			log.ErrorTF(traceID, "DeptToDept %d %d SubDept Fail . Err Is : %v", req.Id, req.ToId, err)
			return baseModel.Fail(constant.DeptToNG)
		}
	}
	go func() { _ = cacheModel.SyncDeptTreeCache(traceID) }()
	return baseModel.Success(constant.DeptToSS, true)
}
