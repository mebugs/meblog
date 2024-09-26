package platService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/cacheModel"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"siteol.com/smart/src/common/utils"
	"siteol.com/smart/src/common/utils/security"
)

// AddAccount 创建登陆账号
func AddAccount(traceID string, req *platModel.AccountAddReq) *baseModel.ResBody {
	// 创建对象初始化
	dbReq, err := req.ToDbReq()
	if err != nil {
		log.ErrorTF(traceID, "AddAccount ToDB Fail . Err Is : %v", err)
		return baseModel.ResFail
	}
	err = platDB.AccountTable.InsertOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "AddAccount Fail . Err Is : %v", err)
		// 解析数据库错误
		return checkAccountDBErr(err)
	}
	// 插入角色关系
	err = syncAccountRoles(traceID, dbReq.Id, req.RoleIds, false)
	if err != nil {
		// 账号角色同步失败
		return baseModel.Fail(constant.AccountRoleNG)
	}
	return baseModel.Success(constant.AccountAddSS, true)
}

// PageAccount 查询登陆账号分页
func PageAccount(traceID string, req *platModel.AccountPageReq) *baseModel.ResBody {
	// 查询分页
	total, list, err := platDB.AccountTable.Page(accountPageQuery(req))
	if err != nil {
		log.ErrorTF(traceID, "PageAccount Fail . Err Is : %v", err)
		return baseModel.Fail(constant.AccountGetNG)
	}
	// 查询这些账号绑定的角色以及部门ID
	rolesLists, deptIds, err := getAccountsRolesAndDept(list)
	if err != nil {
		log.ErrorTF(traceID, "GetAccountsRoles Fail . Err Is : %v", err)
		return baseModel.Fail(constant.AccountGetNG)
	}
	// 查询部门Map
	deptMap, err := getDeptMapByIds(deptIds)
	if err != nil {
		log.ErrorTF(traceID, "GetAccountsDept Fail . Err Is : %v", err)
		return baseModel.Fail(constant.AccountGetNG)
	}
	// 查询角色Map
	roleMap, err := cacheModel.GetRoleCache(traceID)
	if err != nil {
		return baseModel.Fail(constant.AccountGetNG)
	}
	return baseModel.SuccessUnPop(baseModel.SetPageRes(platModel.ToAccountPageRes(list, rolesLists, roleMap, deptMap), total))
}

// GetAccount 登陆账号详情
func GetAccount(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := platDB.AccountTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetAccount Fail . Err Is : %v", err)
		return baseModel.Fail(constant.AccountGetNG)
	}
	// 查询账号绑定的角色
	roleIds, err := platDB.AccountRole{}.GetRoleIds(res.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetAccountRoles Fail . Err Is : %v", err)
		return baseModel.Fail(constant.AccountGetNG)
	}
	// 查询部门Map
	deptMap, err := getDeptMapByIds([]uint64{res.DeptId})
	if err != nil {
		log.ErrorTF(traceID, "GetAccountDept Fail . Err Is : %v", err)
		return baseModel.Fail(constant.AccountGetNG)
	}
	// 查询角色Map
	roleMap, err := cacheModel.GetRoleCache(traceID)
	if err != nil {
		return baseModel.Fail(constant.AccountGetNG)
	}
	return baseModel.SuccessUnPop(platModel.ToAccountGetRes(&res, roleIds, roleMap, deptMap))
}

// EditAccount 编辑登陆账号
func EditAccount(traceID string, req *platModel.AccountEditReq) *baseModel.ResBody {
	dbReq, err := platDB.AccountTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetAccount Fail . Err Is : %v", err)
		return baseModel.Fail(constant.AccountGetNG)
	}
	// 特殊账号禁止编辑
	if dbReq.Mark == constant.StatusLock {
		log.ErrorTF(traceID, "EditAccount %d Fail . Can not Edit", dbReq.Id)
		return baseModel.Fail(constant.AccountMarkNG)
	}
	// 对象更新
	req.ToDbReq(&dbReq)
	err = platDB.AccountTable.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "EditAccount %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return checkAccountDBErr(err)
	}
	// 更新角色关系
	err = syncAccountRoles(traceID, dbReq.Id, req.RoleIds, true)
	if err != nil {
		// 账号角色同步失败
		return baseModel.Fail(constant.AccountRoleNG)
	}
	return baseModel.Success(constant.AccountEditSS, true)
}

// DelAccount 登陆账号移除
func DelAccount(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	dbReq, err := platDB.AccountTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetAccount Fail . Err Is : %v", err)
		return baseModel.Fail(constant.AccountGetNG)
	}
	// 特殊账号禁止刪除
	if dbReq.Mark == constant.StatusLock {
		log.ErrorTF(traceID, "DelAccount %d Fail . Can not Edit", dbReq.Id)
		return baseModel.Fail(constant.AccountMarkNG)
	}
	// 物理删除
	err = platDB.AccountTable.DeleteOne(dbReq.Id)
	if err != nil {
		log.ErrorTF(traceID, "DelAccount %d Fail . Err Is : %v", dbReq.Id, err)
		return baseModel.Fail(constant.AccountDelNG)
	}
	// 更新角色关系
	err = syncAccountRoles(traceID, dbReq.Id, nil, true)
	if err != nil {
		// 账号角色同步失败
		return baseModel.Fail(constant.AccountRoleNG)
	}
	return baseModel.Success(constant.AccountDelSS, true)
}

// ResetAccount 登陆账号重置
func ResetAccount(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	_, err := platDB.AccountTable.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "GetAccount Fail . Err Is : %v", err)
		return baseModel.Fail(constant.AccountGetNG)
	}
	// 重置密码
	saltKey := utils.SaltKey()
	// 初始密码
	pwdC, err := security.AESEncrypt("123456", saltKey)
	if err != nil {
		log.ErrorTF(traceID, "ResetAccount InitPwd Fail . Err Is : %v", err)
		return baseModel.Fail(constant.AccountResetNG)
	}
	// 重置密码
	err = platDB.AccountTable.Executor().ResetAccount(req.Id, saltKey, pwdC, false)
	if err != nil {
		log.ErrorTF(traceID, "ResetAccount %d Fail . Err Is : %v", req.Id, err)
		return baseModel.Fail(constant.AccountResetNG)
	}
	return baseModel.Success(constant.AccountResetSS, true)
}
