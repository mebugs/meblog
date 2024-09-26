package platDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
	"time"
)

// Dept 集团部门
type Dept struct {
	Id             uint64     `json:"id"`             // 默认数据ID
	Name           string     `json:"name"`           // 部门名称
	Pid            uint64     `json:"pid"`            // 父级部门ID，租户创建时默认创建根部门，父级ID=0
	Sort           uint16     `json:"sort"`           // 同级部门排序
	PermissionType string     `json:"permissionType"` // 权限类型，枚举：0_本部门与子部门 1_本部门 2_个人 3_全局 4_指定部门 5_指定人
	Mark           string     `json:"mark"`           // 变更标识，枚举：0_可变更 1_禁止变更
	Status         string     `json:"status"`         // 状态，枚举：0_正常 1_锁定 2_封存
	CreateAt       *time.Time `json:"createAt"`       // 创建时间
	UpdateAt       *time.Time `json:"updateAt"`       // 更新时间
}

// DeptTable 集团部门泛型造器
var DeptTable actuator.Table[Dept]

// DataBase 实现指定数据库
func (t Dept) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t Dept) TableName() string {
	return "dept"
}

// DeptArray 部门自定义排序
type DeptArray []*Dept

func (p DeptArray) Len() int {
	return len(p)
}

func (p DeptArray) Less(i, j int) bool {
	return p[i].Sort < p[j].Sort
}

func (p DeptArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// ToNewDept 迁移到新部门
func (t Dept) ToNewDept(deptId, newDeptId uint64) (err error) {
	now := time.Now()
	r := platDb.Table(t.TableName()).Where("pid = ?", deptId).Updates(map[string]any{
		"pid":       newDeptId,
		"update_at": &now,
	})
	err = r.Error
	return
}
