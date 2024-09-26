package platDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
	"time"
)

// Account 登陆账号
type Account struct {
	Id             uint64     `json:"id"`             // 默认数据ID
	Account        string     `json:"account"`        // 登陆账号
	Encryption     string     `json:"encryption"`     // 密码密文
	SaltKey        string     `json:"saltKey"`        // 密码盐值（AES加密KEY）
	Name           string     `json:"name"`           // 姓名
	DeptId         uint64     `json:"deptId"`         // 部门ID
	IsLeader       string     `json:"isLeader"`       // 部门职位，枚举：0_部门员工 1_部门领导
	PermissionType string     `json:"permissionType"` // 权限类型，枚举：0_继承部门 1_本部门 2_本人 3_全局
	PwdExpTime     *time.Time `json:"pwdExpTime"`     // 密码过期时间，创建即过期
	LastLoginTime  *time.Time `json:"lastLoginTime"`  // 最后登陆时间，为空表示创建
	Mark           string     `json:"mark"`           // 变更标识，枚举：0_可变更 1_禁止变更
	Status         string     `json:"status"`         // 状态，枚举：0_正常 1_锁定 2_封存
	CreateAt       *time.Time `json:"createAt"`       // 创建时间
	UpdateAt       *time.Time `json:"updateAt"`       // 更新时间
}

// AccountTable 登陆账号泛型造器
var AccountTable actuator.Table[Account]

// DataBase 实现指定数据库
func (t Account) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t Account) TableName() string {
	return "account"
}

// ResetAccount 重置账号密码
func (t Account) ResetAccount(id uint64, saltKey, pwdC string, self bool) (err error) {
	now := time.Now()
	exp := now
	if self {
		exp = now.Add(90 * 24 * time.Hour)
	}
	r := platDb.Table(t.TableName()).Where("id = ?", id).Updates(map[string]any{
		"encryption":   pwdC,
		"salt_key":     saltKey,
		"pwd_exp_time": &exp,
		"update_at":    &now,
	})
	err = r.Error
	return
}

// ToNewDept 迁移到新部门
func (t Account) ToNewDept(deptId, newDeptId uint64) (err error) {
	now := time.Now()
	r := platDb.Table(t.TableName()).Where("dept_id = ?", deptId).Updates(map[string]any{
		"dept_id":   newDeptId,
		"update_at": &now,
	})
	err = r.Error
	return
}
