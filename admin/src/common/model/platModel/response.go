package platModel

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"time"
)

// ResponseNextValReq 字典建议Val请求
type ResponseNextValReq struct {
	ServiceCode string `json:"serviceCode" binding:"required,max=3" example:"0"`      // 业务ID，来源于字典，指定响应码归属业务
	Type        string `json:"type" binding:"required,oneof='S' 'F' 'E'" example:"S"` // 响应类型，该字段用于筛选，可配置S/F/E
}

// ResponseDoReq 响应码通用请求
type ResponseDoReq struct {
	ZhCn   string `json:"zhCn" binding:"required,max=128" example:"ZhCn"` // 中文响应文言
	EnUs   string `json:"enUs" binding:"required,max=128" example:"EnUs"` // 英文响应文言
	Remark string `json:"remark" binding:"max=128" example:"Remark"`      // 其他备注信息
}

// ResponseAddReq 响应码创建请求
type ResponseAddReq struct {
	ResponseDoReq
	ServiceCode string `json:"serviceCode" binding:"required,max=3" example:"0"`      // 业务ID，来源于字典，指定响应码归属业务
	Type        string `json:"type" binding:"required,oneof='S' 'F' 'E'" example:"S"` // 响应类型，该字段用于筛选，可配置S/F/E
	Code        string `json:"code" example:"S101"`                                   // 响应码，仅示例，实际入库实时计算
}

// ToDbReq 响应码创建对象转字典对象
func (r *ResponseAddReq) ToDbReq() *platDB.ResponseCode {
	now := time.Now()
	return &platDB.ResponseCode{
		Id:          0,
		ServiceCode: r.ServiceCode,
		Type:        r.Type,
		ZhCn:        r.ZhCn,
		EnUs:        r.EnUs,
		Remark:      r.Remark,
		Common: platDB.Common{
			Mark:     constant.StatusOpen,
			Status:   constant.StatusOpen,
			CreateAt: &now,
			UpdateAt: &now,
		},
	}
}

// ResponseEditReq 响应码编辑请求
type ResponseEditReq struct {
	Id uint64 `json:"id" binding:"required" example:"1"` // 数据ID
	ResponseDoReq
}

// ToDbReq 响应码更新对象转字典对象
func (r *ResponseEditReq) ToDbReq(d *platDB.ResponseCode) {
	now := time.Now()
	d.ZhCn = r.ZhCn
	d.EnUs = r.EnUs
	d.Remark = r.Remark
	d.UpdateAt = &now
}

// ResponseGetRes 响应码分页响应
type ResponseGetRes struct {
	Id          uint64 `json:"id" example:"1"`          // 数据ID
	Code        string `json:"code" example:"20101"`    // 响应码
	ServiceCode string `json:"serviceCode" example:"1"` // 业务ID，来源于字典，指定响应码归属业务
	Type        string `json:"type" example:"S"`        // 响应类型，该字段用于筛选，可配置S/F/E
	ZhCn        string `json:"zhCn" example:"ZhCn"`     // 中文响应文言
	EnUs        string `json:"enUs" example:"EnUs"`     // 英文响应文言
	Remark      string `json:"remark" example:"Remark"` // 其他备注信息
}

// ToResponseGetRes 响应码转为查询对象
func ToResponseGetRes(r *platDB.ResponseCode) *ResponseGetRes {
	return &ResponseGetRes{
		Id:          r.Id,
		Code:        r.Code,
		ServiceCode: r.ServiceCode,
		Type:        r.Type,
		ZhCn:        r.ZhCn,
		EnUs:        r.EnUs,
		Remark:      r.Remark,
	}
}

// ResponsePageReq 响应码分页请求
type ResponsePageReq struct {
	Code        string `json:"code" example:"20"`      // 响应码，支持模糊查询
	ServiceCode string `json:"serviceCode"example:"0"` // 业务ID，来源于字典，指定响应码归属业务
	Type        string `json:"type" example:"S"`       // 响应类型，该字段用于筛选，可配置S/F/E
	baseModel.PageReq
}

// ResponsePageRes 响应码分页响应
type ResponsePageRes struct {
	ResponseGetRes
	Mark string `json:"mark" example:"1"` // 变更标识 0可变更 1禁止变更
}

// ToResponsePageRes 响应码转为分页对象
func ToResponsePageRes(list []*platDB.ResponseCode) []*ResponsePageRes {
	res := make([]*ResponsePageRes, len(list))
	for i, r := range list {
		res[i] = &ResponsePageRes{
			ResponseGetRes: *ToResponseGetRes(r),
			Mark:           r.Mark,
		}
	}
	return res
}
