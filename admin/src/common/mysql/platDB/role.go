package platDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
)

// Role 角色
type Role struct {
	Id     uint64 `json:"id"`     // 默认数据ID
	Name   string `json:"name"`   // 角色名称
	Remark string `json:"remark"` // 角色备注
	Common
}

// RoleTable 角色泛型造器
var RoleTable actuator.Table[Role]

// DataBase 实现指定数据库
func (t Role) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t Role) TableName() string {
	return "role"
}
