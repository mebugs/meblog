package platModel

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"time"
)

// RouterDoReq 路由接口通用请求
type RouterDoReq struct {
	Name         string `json:"name" binding:"required,max=32" example:"Login"`                  // 路由名称，用于界面展示，与权限关联
	LogInDb      string `json:"logInDb" binding:"required,oneof='0' '1'" example:"0"`            // 日志入库 0 启用 1 默认不启用
	ReqLogPrint  string `json:"reqLogPrint" binding:"required,oneof='0' '1'" example:"0"`        // 请求日志打印 0 打印 1 不打印
	ReqLogSecure string `json:"reqLogSecure" binding:"max=256" example:"phone,account,password"` // 请求日志脱敏字段，逗号分隔，打印时允许配置
	ResLogPrint  string `json:"resLogPrint" binding:"required,oneof='0' '1'" example:"0"`        // 响应日志打印 0 打印 1 不打印
	ResLogSecure string `json:"resLogSecure" binding:"max=256" example:"name,account,password"`  // 响应日志脱敏字段，逗号分隔，打印时允许配置
	Remark       string `json:"remark" binding:"max=128" example:"login"`                        // 其他备注信息
}

// RouterAddReq 路由接口创建请求
type RouterAddReq struct {
	RouterDoReq
	Type        string `json:"type" binding:"required,oneof='0' '1'" example:"0"`       // 免授权路由 0 免授权 1 授权
	ServiceCode string `json:"serviceCode" binding:"required,max=3" example:"base"`     // 业务编码（字典），为接口分组
	Url         string `json:"url" binding:"required,uri,max=64" example:"/auth/login"` // 路由地址，后端访问URL，后端不在URL中携带参数，统一Post处理内容
}

// ToDbReq 路由接口创建对象转字典对象
func (r *RouterAddReq) ToDbReq() *platDB.Router {
	now := time.Now()
	return &platDB.Router{
		Id:           0,
		Name:         r.Name,
		Url:          r.Url,
		Type:         r.Type,
		ServiceCode:  r.ServiceCode,
		LogInDb:      r.LogInDb,
		ReqLogPrint:  r.ReqLogPrint,
		ReqLogSecure: r.ReqLogSecure,
		ResLogPrint:  r.ResLogPrint,
		ResLogSecure: r.ResLogSecure,
		Remark:       r.Remark,
		Common: platDB.Common{
			Mark:     constant.StatusOpen,
			Status:   constant.StatusOpen,
			CreateAt: &now,
			UpdateAt: &now,
		},
	}
}

// RouterEditReq 路由接口编辑请求
type RouterEditReq struct {
	Id uint64 `json:"id" binding:"required" example:"1"` // 数据ID
	RouterDoReq
}

// ToDbReq 路由接口更新对象转字典对象
func (r *RouterEditReq) ToDbReq(d *platDB.Router) {
	now := time.Now()
	d.Name = r.Name
	d.LogInDb = r.LogInDb
	d.ReqLogPrint = r.ReqLogPrint
	d.ReqLogSecure = r.ReqLogSecure
	d.ResLogPrint = r.ResLogPrint
	d.ResLogSecure = r.ResLogSecure
	d.Remark = r.Remark
	d.UpdateAt = &now
}

// RouterGetRes 路由接口分页响应
type RouterGetRes struct {
	Id           uint64 `json:"id" example:"1"`               // 数据ID
	Type         string `json:"type" example:"0"`             // 免授权路由 0 免授权 1 授权
	Name         string `json:"name" example:"Login"`         // 路由名称，用于界面展示，与权限关联
	Url          string `json:"url" example:"/auth/login"`    // 路由地址，后端访问URL，后端不在URL中携带参数，统一Post处理内容
	ServiceCode  string `json:"serviceCode" example:"base"`   // 业务编码（字典），为接口分组
	LogInDb      string `json:"logInDb" example:"0"`          // 日志入库 0 启用 1 默认不启用
	ReqLogPrint  string `json:"reqLogPrint" example:"0"`      // 请求日志打印 0 打印 1 不打印
	ReqLogSecure string `json:"reqLogSecure" example:"phone"` // 请求日志脱敏字段，逗号分隔，打印时允许配置
	ResLogPrint  string `json:"resLogPrint" example:"0"`      // 响应日志打印 0 打印 1 不打印
	ResLogSecure string `json:"resLogSecure" example:"name"`  // 响应日志脱敏字段，逗号分隔，打印时允许配置
	Remark       string `json:"remark" example:"login"`       // 其他备注信息
}

// ToRouterGetRes 路由接口转为查询对象
func ToRouterGetRes(r *platDB.Router) *RouterGetRes {
	return &RouterGetRes{
		Id:           r.Id,
		Type:         r.Type,
		Name:         r.Name,
		Url:          r.Url,
		ServiceCode:  r.ServiceCode,
		LogInDb:      r.LogInDb,
		ReqLogPrint:  r.ReqLogPrint,
		ReqLogSecure: r.ReqLogSecure,
		ResLogPrint:  r.ResLogPrint,
		ResLogSecure: r.ResLogSecure,
		Remark:       r.Remark,
	}
}

// RouterPageReq 路由接口分页请求
type RouterPageReq struct {
	Name        string `json:"name" example:"login"`        // 路由名，支持模糊查询
	Url         string `json:"url" example:"/auth/login"`   // 路由地址，后端访问URL，支持模糊查询
	ServiceCode string `json:"serviceCode"  example:"base"` // 业务编码（字典），为接口分组
	Type        string `json:"type" example:"0"`            // 免授权路由 0 免授权 1 授权
	baseModel.PageReq
}

// RouterPageRes 路由接口分页响应
type RouterPageRes struct {
	RouterGetRes
	Mark string `json:"mark" example:"1"` // 变更标识 0可变更 1禁止变更
}

// ToRouterPageRes 路由接口转为分页对象
func ToRouterPageRes(list []*platDB.Router) []*RouterPageRes {
	res := make([]*RouterPageRes, len(list))
	for i, r := range list {
		res[i] = &RouterPageRes{
			RouterGetRes: *ToRouterGetRes(r),
			Mark:         r.Mark,
		}
	}
	return res
}
