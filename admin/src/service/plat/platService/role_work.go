package platService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/common/mysql/actuator"
	"strings"
)

// 业务层数据处理函数
// 抽取到独立文件中仅便于Server层阅读（没有特别意义）

// 解析数据库错误
func checkRoleDBErr(err error) *baseModel.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		if strings.Contains(errStr, "name_uni") {
			// 角色名唯一
			return baseModel.Fail(constant.RoleNameUniNG)
		}
	}
	// 默认业务异常
	return baseModel.ResFail
}

// 字典分页查询对象
func rolePageQuery(req *platModel.RolePageReq) (query *actuator.Query) {
	// 初始化Page
	req.PageReq.PageInit()
	// 组装Query
	query = actuator.InitQuery()
	if req.Name != "" {
		query.Like("name", req.Name)
	}
	query.Eq("status", constant.StatusOpen)
	query.Desc("id")
	query.LimitByPage(req.Current, req.PageSize)
	return
}
