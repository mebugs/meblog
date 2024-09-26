package platDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
)

// ResponseCode 响应码配置
type ResponseCode struct {
	Id          uint64 `json:"id"`          // 数据ID
	Code        string `json:"code"`        // 响应码
	ServiceCode string `json:"serviceCode"` // 业务ID，来源于字典，指定响应码归属业务
	Type        string `json:"type"`        // 响应类型，该字段用于筛选，可配置S和F
	ZhCn        string `json:"zhCn"`        // 中文响应文言
	EnUs        string `json:"enUs"`        // 英文响应文言
	Remark      string `json:"remark"`      // 其他备注信息
	Common
}

// ResponseTable 响应码配置泛型造器
var ResponseTable actuator.Table[ResponseCode]

// DataBase 实现指定数据库
func (t ResponseCode) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t ResponseCode) TableName() string {
	return "response_code"
}
