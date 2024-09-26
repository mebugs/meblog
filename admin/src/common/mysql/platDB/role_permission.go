package platDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
)

// RolePermission 角色权限
type RolePermission struct {
	Id           uint64 `json:"id"`           // 默认数据ID
	RoleId       uint64 `json:"roleId"`       // 角色ID
	PermissionId uint64 `json:"permissionId"` // 权限ID
	CheckType    string `json:"checkType"`    // 选择类型 0 check 1 halfCheck
}

// RolePermissionTable 角色权限泛型造器
var RolePermissionTable actuator.Table[RolePermission]

// DataBase 实现指定数据库
func (t RolePermission) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t RolePermission) TableName() string {
	return "role_permission"
}

// GetPermissionIdsWithRoleIds 获取角色IDS的权限ID集
func (t RolePermission) GetPermissionIdsWithRoleIds(roleIds []uint64) (res []uint64, err error) {
	r := platDb.Table(t.TableName()).Distinct("permission_id").Where("role_id IN ? ", roleIds).Find(&res)
	err = r.Error
	return
}

// GetRoleIds 获取权益ID对应的角色
func (t RolePermission) GetRoleIds(permissionId uint64) (res []uint64, err error) {
	r := platDb.Table(t.TableName()).Distinct("role_id").Where("permission_id = ?", permissionId).Find(&res)
	err = r.Error
	return
}

// DeleteByRoleId 根据路由ID移除权限
func (t RolePermission) DeleteByRoleId(roleId uint64) (err error) {
	r := platDb.Where("role_id = ?", roleId).Delete(&t)
	err = r.Error
	return
}

// DeleteByPermissionId 根据权限ID移除路由
func (t RolePermission) DeleteByPermissionId(permissionId uint64) (err error) {
	r := platDb.Where("permission_id = ?", permissionId).Delete(&t)
	err = r.Error
	return
}
