package blogService

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/common/mysql/actuator"
	"strings"
)

// 业务层数据处理函数
// 抽取到独立文件中仅便于Server层阅读（没有特别意义）

// 解析数据库错误
func checkCategoryDBErr(err error) *baseModel.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		if strings.Contains(errStr, "url_uni") {
			// 唯一索引错误
			return baseModel.Fail(constant.CategoryUniXxxNG)
		}
	}
	// 默认业务异常
	return baseModel.ResFail
}

// 分页查询对象封装
func categoryPageQuery(req *blogModel.CategoryPageReq) (query *actuator.Query) {
	// 初始化Page
	req.PageReq.PageInit()
	// 组装Query
	query = actuator.InitQuery()
	// 模拟代码，更多函数参考Query构造器
	if req.Id != 0 {
		query.Like("id", req.Id)
	}
	if req.Id != 0 {
		query.Eq("id", req.Id)
	}
	// 模拟代码，更多函数参考Query构造器
	query.Eq("status", constant.StatusOpen)
	query.Desc("id")
	query.LimitByPage(req.Current, req.PageSize)
	return
}
