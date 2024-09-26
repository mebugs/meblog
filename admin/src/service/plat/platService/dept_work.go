package platService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"strings"
)

// 业务层数据处理函数
// 抽取到独立文件中仅便于Server层阅读（没有特别意义）

// 解析数据库错误
func checkDeptDBErr(err error) *baseModel.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		if strings.Contains(errStr, "dept_uni") {
			// 唯一索引错误
			return baseModel.Fail(constant.DeptUniNameNG)
		}
	}
	// 默认业务异常
	return baseModel.ResFail
}

// getDeptMapByIds 根据ID获取部门MAP
func getDeptMapByIds(ids []uint64) (deptMap map[uint64]string, err error) {
	if len(ids) == 0 {
		return
	}
	deptS, err := platDB.DeptTable.GetByIds(ids)
	if err != nil {
		return
	}
	deptMap = make(map[uint64]string, len(deptS))
	for _, item := range deptS {
		deptMap[item.Id] = item.Name
	}
	return
}

// getDeptAccounts 部门账号列表
func getDeptAccounts(id uint64) (res [][]*platModel.DeptAccountRes, err error) {
	// 查询部门成员
	accounts, err := platDB.AccountTable.GetByObject(&platDB.Account{DeptId: id})
	if err != nil {
		return
	}
	res = make([][]*platModel.DeptAccountRes, 2)
	for i := range res {
		res[i] = make([]*platModel.DeptAccountRes, 0)
	}
	for _, item := range accounts {
		// 对于部门领导大于第一组
		resItem := &platModel.DeptAccountRes{
			Id:      item.Id,
			Name:    item.Name,
			Account: item.Account,
		}
		if item.IsLeader == constant.StatusLock {
			res[0] = append(res[0], resItem)
		} else {
			res[1] = append(res[1], resItem)
		}
	}
	return
}
