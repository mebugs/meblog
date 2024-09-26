package platModel

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"time"
)

// DictReadReq 字典读取请求
type DictReadReq struct {
	GroupKeys []string `json:"groupKeys" binding:"required" example:"serviceCode,responseType"` // 需要查询的字典分组
	Local     string   `json:"-"`                                                               // 字典语言
}

// DictReadRes 字典读取响应
type DictReadRes struct {
	List map[string][]*baseModel.SelectRes `json:"list"` // 字典下拉列表 {'serviceCode':"[{'label':'基础','value':'1'}]"}
	Map  map[string]map[string]string      `json:"map"`  // 字典翻译Map {'serviceCode':{'1':'基础'}}
}

// DictNextValReq 字典建议Val请求
type DictNextValReq struct {
	GroupKey string `json:"groupKey" binding:"required" example:"serviceCode"` // 字典分组KEY
}

// DictBroReq 字典分组列表请求
type DictBroReq struct {
	GroupKey string `json:"groupKey" binding:"required" example:"serviceCode"` // 字典分组KEY
	Local    string `json:"-"`                                                 // 字典语言
}

// DictDoReq 字典处理请求（共通）
type DictDoReq struct {
	Label   string `json:"label" binding:"required,max=64" example:"PlatBase"`          // 字段名称
	LabelEn string `json:"labelEn" binding:"required,max=64" example:"PlatBase"`        // 字段名称（英文）
	Choose  string `json:"choose" binding:"required,oneof='0' '1'" example:"0"`         // 是否可被选择 0可选择 1不可选择
	Remark  string `json:"remark" binding:"max=128" example:"Business type dictionary"` // 字典描述
	// Pid     uint64 `json:"pid"` // 父级字典ID 默认 1（根数据），暂时不开放变更
}

// DictAddReq 字典创建请求
type DictAddReq struct {
	GroupKey  string `json:"groupKey" binding:"required,max=32" example:"serviceCode"` // 字典分组KEY
	Val       string `json:"val" binding:"required,max=32" example:"1"`                // 字典值（字符型）
	DictDoReq        // 共通引用
}

// ToDbReq 字典创建对象转字典对象
func (r *DictAddReq) ToDbReq() *platDB.Dict {
	now := time.Now()
	return &platDB.Dict{
		Id:       0,
		GroupKey: r.GroupKey,
		Label:    r.Label,
		LabelEn:  r.LabelEn,
		Choose:   r.Choose,
		Val:      r.Val,
		Pid:      1,
		Sort:     0,
		Remark:   r.Remark,
		Common: platDB.Common{
			Mark:     constant.StatusOpen,
			Status:   constant.StatusOpen,
			CreateAt: &now,
			UpdateAt: &now,
		},
	}
}

// DictEditReq 字典编辑请求
type DictEditReq struct {
	Id        uint64 `json:"id" binding:"required" example:"1"` // 数据ID
	DictDoReq        // 共通引用
}

// ToDbReq 字典更新对象转字典对象
func (r *DictEditReq) ToDbReq(d *platDB.Dict) {
	now := time.Now()
	d.Label = r.Label
	d.LabelEn = r.LabelEn
	d.Choose = r.Choose
	d.Remark = r.Remark
	d.UpdateAt = &now
}

// DictGetRes 字典查询结果
type DictGetRes struct {
	Id       uint64 `json:"id" example:"1"`                 // 数据ID
	GroupKey string `json:"groupKey" example:"serviceCode"` // 字典分组KEY
	Val      string `json:"val" example:"1"`                // 字典值（字符型）
	Label    string `json:"label" example:"PlatBase"`       // 字段名称
	LabelEn  string `json:"labelEn" example:"PlatBase"`     // 字段名称（英文）
	Choose   string `json:"choose" example:"0"`             // 是否可被选择 0可选择 1不可选择
	Remark   string `json:"remark" example:"Remark"`        // 字典描述
}

// ToDictGetRes 字典转为查询对象
func ToDictGetRes(r *platDB.Dict) *DictGetRes {
	return &DictGetRes{
		Id:       r.Id,
		GroupKey: r.GroupKey,
		Val:      r.Val,
		Label:    r.Label,
		LabelEn:  r.LabelEn,
		Choose:   r.Choose,
		Remark:   r.Remark,
	}
}

// DictPageReq 字典分页请求
type DictPageReq struct {
	GroupKey string `json:"groupKey" example:"serviceCode"` // 需要查询的字典分组
	Choose   string `json:"choose" example:"1"`             // 字典是否可选择
	baseModel.PageReq
}

// DictPageRes 字典分页响应
type DictPageRes struct {
	DictGetRes
	Sort uint16 `json:"sort" example:"0"` // 字典排序
	Mark string `json:"mark" example:"1"` // 变更标识 0可变更 1禁止变更
}

// ToDictPageRes 字典转为分页对象
func ToDictPageRes(list []*platDB.Dict) []*DictPageRes {
	res := make([]*DictPageRes, len(list))
	for i, r := range list {
		res[i] = &DictPageRes{
			DictGetRes: *ToDictGetRes(r),
			Sort:       r.Sort,
			Mark:       r.Mark,
		}
	}
	return res
}
