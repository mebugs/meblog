package platDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
	"time"
)

// RouterLog 路由接口日志
type RouterLog struct {
	Id         uint64     `json:"id"`         // 数据ID
	AppName    string     `json:"appName"`    // 应用名
	AppNode    string     `json:"appNode"`    // 应用节点
	AppTraceId string     `json:"appTraceId"` // 应用节点TraceID
	ReqIp      string     `json:"reqIp"`      // 来源IP
	ReqUrl     string     `json:"reqUrl"`     // 请求地址
	ReqBody    string     `json:"reqBody"`    // 请求报文
	ReqAt      *time.Time `json:"reqAt"`      // 请求时间
	ResStatus  int        `json:"resStatus"`  // 响应状态
	ResBody    string     `json:"resBody"`    // 响应报文
	ResTime    *time.Time `json:"resTime"`    // 响应时间
}

// RouterLogTable 路由接口日志泛型造器
var RouterLogTable actuator.Table[RouterLog]

// DataBase 实现指定数据库
func (t RouterLog) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t RouterLog) TableName() string {
	return "router_log"
}
