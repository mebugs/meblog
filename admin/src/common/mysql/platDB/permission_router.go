package platDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
)

// PermissionRouter 权限路由
type PermissionRouter struct {
	Id           uint64 `json:"id"`           // 默认数据ID
	PermissionId uint64 `json:"permissionId"` // 权限ID
	RouterId     uint64 `json:"routerId"`     // 路由ID
}

// PermissionRouterTable 权限路由泛型造器
var PermissionRouterTable actuator.Table[PermissionRouter]

// DataBase 实现指定数据库
func (t PermissionRouter) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t PermissionRouter) TableName() string {
	return "permission_router"
}

// GetRouterIds 获取权限对应的路由ID
func (t PermissionRouter) GetRouterIds(permissionId uint64) (res []uint64, err error) {
	r := platDb.Table(t.TableName()).Distinct("router_id").Where("permission_id = ?", permissionId).Find(&res)
	err = r.Error
	return
}

// GetRouterIdsWithByPermissionIds 获取权限对应的路由ID
func (t PermissionRouter) GetRouterIdsWithByPermissionIds(ids []uint64) (res []uint64, err error) {
	r := platDb.Table(t.TableName()).Distinct("router_id").Where("permission_id IN ?", ids).Find(&res)
	err = r.Error
	return
}

// DeleteByPermissionId 根据权限ID移除路由
func (t PermissionRouter) DeleteByPermissionId(permissionId uint64) (err error) {
	r := platDb.Where("permission_id = ?", permissionId).Delete(&t)
	err = r.Error
	return
}

// DeleteByRouterId 根据路由ID移除权限
func (t PermissionRouter) DeleteByRouterId(routerId uint64) (err error) {
	r := platDb.Where("router_id = ?", routerId).Delete(&t)
	err = r.Error
	return
}
