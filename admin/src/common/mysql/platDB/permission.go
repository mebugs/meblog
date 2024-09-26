package platDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
)

// Permission 权限
type Permission struct {
	Id     uint64 `json:"id"`     // 默认数据ID
	Name   string `json:"name"`   // 权限名称，界面展示，建议与界面导航一致
	Alias  string `json:"alias"`  // 权限别名，英文，规范如下：sys，sysAccount sysAccountAdd
	Level  string `json:"level"`  // 权限等级 1分组（一级导航）2模块（页面）3功能（按钮）第四级路由不在本表中体现
	Pid    uint64 `json:"pid"`    // 父级ID，默认为1
	Sort   uint16 `json:"sort"`   // 字典排序
	Static string `json:"static"` // 默认启用权限，0 启用 1 不启，启用后，该权限默认被分配，不可去勾
	Common
}

// PermissionTable 权限泛型造器
var PermissionTable actuator.Table[Permission]

// DataBase 实现指定数据库
func (t Permission) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t Permission) TableName() string {
	return "permission"
}

// PermissionArray 权限自定义排序
type PermissionArray []*Permission

func (p PermissionArray) Len() int {
	return len(p)
}

func (p PermissionArray) Less(i, j int) bool {
	return p[i].Sort < p[j].Sort
}

func (p PermissionArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
