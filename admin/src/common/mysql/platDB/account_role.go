package platDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
)

// AccountRole 账号角色
type AccountRole struct {
	Id        uint64 `json:"id"`        // 默认数据ID
	AccountId uint64 `json:"accountId"` // 账号ID
	RoleId    uint64 `json:"roleId"`    // 角色ID
}

// AccountRoleTable 账号角色泛型造器
var AccountRoleTable actuator.Table[AccountRole]

// DataBase 实现指定数据库
func (t AccountRole) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t AccountRole) TableName() string {
	return "account_role"
}

// GetRoleIds 获取权限对应的路由ID
func (t AccountRole) GetRoleIds(accountId uint64) (res []uint64, err error) {
	r := platDb.Table(t.TableName()).Distinct("role_id").Where("account_id = ?", accountId).Find(&res)
	err = r.Error
	return
}

// GetAccountIds 获取角色ID对应的账号
func (t AccountRole) GetAccountIds(roleId uint64) (res []uint64, err error) {
	r := platDb.Table(t.TableName()).Distinct("account_id").Where("role_id = ?", roleId).Find(&res)
	err = r.Error
	return
}

// GetAccountIdsWithRoleIds 获取角色IDS的账号ID集
func (t AccountRole) GetAccountIdsWithRoleIds(roleIds []uint64) (res []uint64, err error) {
	r := platDb.Table(t.TableName()).Distinct("account_id").Where("role_id IN ? ", roleIds).Find(&res)
	err = r.Error
	return
}

// DeleteByAccountId 根据账号ID删除角色
func (t AccountRole) DeleteByAccountId(accountId uint64) (err error) {
	r := platDb.Where("account_id = ?", accountId).Delete(&t)
	err = r.Error
	return
}

// DeleteByRoleId 根据角色ID删除账号
func (t AccountRole) DeleteByRoleId(roleId uint64) (err error) {
	r := platDb.Where("role_id = ?", roleId).Delete(&t)
	err = r.Error
	return
}
