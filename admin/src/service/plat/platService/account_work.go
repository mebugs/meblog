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
func checkAccountDBErr(err error) *baseModel.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		if strings.Contains(errStr, "xxx_uni") {
			// 唯一索引错误
			return baseModel.Fail(constant.AccountUniXxxNG)
		}
	}
	// 默认业务异常
	return baseModel.ResFail
}

// 分页查询对象封装
func accountPageQuery(req *platModel.AccountPageReq) (query *actuator.Query) {
	// 初始化Page
	req.PageReq.PageInit()
	// 组装Query
	query = actuator.InitQuery()
	if req.Account != "" {
		query.Like("account", req.Account)
	}
	if req.Name != "" {
		query.Like("name", req.Name)
	}
	if req.DeptId != "" {
		query.Eq("dept_id", req.DeptId)
	}
	query.Desc("id")
	query.LimitByPage(req.Current, req.PageSize)
	return
}
