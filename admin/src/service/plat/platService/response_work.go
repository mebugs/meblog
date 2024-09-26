package platService

import (
	"fmt"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/common/mysql/actuator"
	"siteol.com/smart/src/common/mysql/platDB"
	"strings"
)

// 业务层数据处理函数
// 抽取到独立文件中仅便于Server层阅读（没有特别意义）

// 解析数据库错误
func checkResponseDBErr(err error) *baseModel.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		if strings.Contains(errStr, "code_uni") {
			// 响应码分组下响应码值唯一
			return baseModel.Fail(constant.ResponseUniNG)
		}
	}
	// 默认业务异常
	return baseModel.ResFail
}

// 生成下一个响应码Code
func responseMakeCode(serviceCode, codeType string) (code string, err error) {
	// 组装Query
	query := actuator.InitQuery()
	query.Eq("service_code", serviceCode)
	query.Eq("type", codeType)
	total, err := platDB.ResponseTable.CountByQuery(query)
	if err != nil {
		return
	}
	code = fmt.Sprintf(constant.CodeFmt, codeType, serviceCode, total)
	return
}

// 字典分页查询对象
func responsePageQuery(req *platModel.ResponsePageReq) (query *actuator.Query) {
	// 初始化Page
	req.PageReq.PageInit()
	// 组装Query
	query = actuator.InitQuery()
	if req.ServiceCode != "" {
		query.Eq("service_code", req.ServiceCode)
	}
	if req.Type != "" {
		query.Eq("type", req.Type)
	}
	if req.Code != "" {
		query.Like("code", req.Code)
	}
	query.Eq("status", constant.StatusOpen)
	query.Desc("id")
	query.LimitByPage(req.Current, req.PageSize)
	return
}
