package platDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
)

// Router 接口路由
type Router struct {
	Id           uint64 `json:"id"`           // 默认数据ID
	Name         string `json:"name"`         // 路由名称，用于界面展示，与权限关联
	Url          string `json:"url"`          // 路由地址，后端访问URL，后端不在URL中携带参数，统一Post处理内容
	Type         string `json:"type"`         // 免授权路由 0 免授权 1 授权
	ServiceCode  string `json:"serviceCode"`  // 业务编码（字典），为接口分组
	LogInDb      string `json:"logInDb"`      // 请日志入库 0 打印 1 不打印
	ReqLogPrint  string `json:"reqLogPrint"`  // 请求日志打印 0 打印 1 不打印
	ReqLogSecure string `json:"reqLogSecure"` // 请求日志脱敏字段，逗号分隔，打印时允许配置
	ResLogPrint  string `json:"resLogPrint"`  // 响应日志打印 0 打印 1 不打印
	ResLogSecure string `json:"resLogSecure"` // 响应日志脱敏字段，逗号分隔，打印时允许配置
	Remark       string `json:"remark"`       // 其他备注信息
	Common
}

// RouterTable 接口路由泛型造器
var RouterTable actuator.Table[Router]

// DataBase 实现指定数据库
func (t Router) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t Router) TableName() string {
	return "router"
}
