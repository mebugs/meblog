package platService

import (
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
func dictCheckDBErr(err error) *baseModel.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		if strings.Contains(errStr, "dict_uni") {
			// 字典分组下字典值唯一
			return baseModel.Fail(constant.DictUniNG)
		}
	}
	// 默认业务异常
	return baseModel.ResFail
}

// 字典分页查询对象
func dictPageQuery(req *platModel.DictPageReq) (query *actuator.Query) {
	// 初始化Page
	req.PageReq.PageInit()
	// 组装Query
	query = actuator.InitQuery()
	if req.Choose != "" {
		query.Eq("choose", req.Choose)
	}
	if req.GroupKey != "" {
		query.Eq("group_key", req.GroupKey)
		query.Asc("sort")
	} else {
		query.Desc("id")
	}
	query.Eq("status", constant.StatusOpen)
	query.LimitByPage(req.Current, req.PageSize)
	return
}

// 数据库查询列表转换选择列表与翻译列表
func dictListToBro(list []*platDB.Dict, local string) []*baseModel.SortRes {
	broList := make([]*baseModel.SortRes, 0)
	for _, dict := range list {
		// 未启用字典
		if dict.Status != constant.StatusOpen {
			continue
		}
		// 如果可选择，创建选择数据
		// 翻译label
		var label string
		switch local {
		case "en":
			label = dict.LabelEn
		default:
			label = dict.Label
		}
		broList = append(broList, &baseModel.SortRes{
			Id:   dict.Id,
			Name: label,
			Sort: dict.Sort,
		})
	}
	return broList
}

// 数据库查询列表转换选择列表与翻译列表
func dictListToRead(list []*platDB.Dict, local string) ([]*baseModel.SelectRes, map[string]string) {
	labelList := make([]*baseModel.SelectRes, 0)
	valueMap := make(map[string]string, 0)
	for _, dict := range list {
		// 翻译label
		var label string
		switch local {
		case "en":
			label = dict.LabelEn
		default:
			label = dict.Label
		}
		// 创建Map数据
		valueMap[dict.Val] = label
		// 未启用字典（可以翻译，不可以选择）
		if dict.Status != constant.StatusOpen {
			continue
		}
		// 如果可选择，创建选择数据
		if dict.Choose == constant.StatusOpen {
			labelList = append(labelList, &baseModel.SelectRes{Label: label, Value: dict.Val})
		}
	}
	return labelList, valueMap
}
