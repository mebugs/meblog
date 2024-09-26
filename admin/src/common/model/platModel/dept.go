package platModel

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"time"
)

// DeptDoReq 集团部门 通用请求，创建&编辑可复用的字段
type DeptDoReq struct {
	Name           string `json:"name" binding:"required,max=16" example:"demo"`                               // 部门名称
	Pid            uint64 `json:"pid" binding:"required" example:"0"`                                          // 父级部门ID，租户创建时默认创建根部门，父级ID=0
	PermissionType string `json:"permissionType" binding:"required,oneof='0' '1' '2' '3' '4' '5'" example:"0"` // 权限类型，枚举：0_本部门与子部门 1_本部门 2_个人 3_全局 4_指定部门 5_指定人
}

// DeptAddReq 集团部门 创建请求，酌情从通用中摘出部分字段
type DeptAddReq struct {
	DeptDoReq
}

// DeptEditReq 集团部门 编辑请求，酌情从通用中摘出部分字段
type DeptEditReq struct {
	Id uint64 `json:"id" binding:"required" example:"1"` // 数据ID
	DeptDoReq
}

// ToDbReq 集团部门 创建转数据库
func (r *DeptAddReq) ToDbReq() *platDB.Dept {
	now := time.Now()
	return &platDB.Dept{
		Id:             0,
		Name:           r.Name,
		Pid:            r.Pid,
		Sort:           0,
		PermissionType: r.PermissionType,
		Mark:           constant.StatusOpen,
		Status:         constant.StatusOpen,
		CreateAt:       &now,
		UpdateAt:       &now,
	}
}

// ToDbReq 集团部门 更新转数据库
func (r *DeptEditReq) ToDbReq(d *platDB.Dept) {
	now := time.Now()
	d.Name = r.Name
	d.PermissionType = r.PermissionType
	d.UpdateAt = &now
}

// DeptGetRes 集团部门 详情响应
type DeptGetRes struct {
	Id             uint64              `json:"id" example:"1"`             // 数据ID
	Name           string              `json:"name" example:"demo"`        // 部门名称
	Pid            uint64              `json:"pid" example:"0"`            // 父级部门ID，租户创建时默认创建根部门，父级ID=0
	Sort           uint16              `json:"sort" example:"0"`           // 同级部门排序
	PermissionType string              `json:"permissionType" example:"0"` // 权限类型，枚举：0_本部门与子部门 1_本部门 2_个人 3_指定部门 4_指定人 5_全局
	Mark           string              `json:"mark" example:"0"`           // 变更标识，枚举：0_可变更 1_禁止变更
	Accounts       [][]*DeptAccountRes `json:"accounts"`                   // 部门账号，第1组是领导，第二组是成员
}

// DeptAccountRes 集团部门 账号响应
type DeptAccountRes struct {
	Id      uint64 `json:"id" example:"1"`         // 数据ID
	Account string `json:"account" example:"demo"` // 账号
	Name    string `json:"name" example:"demo"`    // 姓名
}

// ToDeptGetRes 集团部门 数据库转为详情响应
func ToDeptGetRes(r *platDB.Dept, accounts [][]*DeptAccountRes) *DeptGetRes {
	return &DeptGetRes{
		Id:             r.Id,
		Name:           r.Name,
		Pid:            r.Pid,
		Sort:           r.Sort,
		PermissionType: r.PermissionType,
		Mark:           r.Mark,
		Accounts:       accounts,
	}
}

// ToDeptBroRes 部门转为查询对象
func ToDeptBroRes(r platDB.DeptArray) []*baseModel.SortRes {
	res := make([]*baseModel.SortRes, len(r))
	for i, item := range r {
		res[i] = &baseModel.SortRes{
			Id:   item.Id,
			Name: item.Name,
			Sort: item.Sort,
		}
	}
	return res
}

// DeptToReq 集团部门 迁移请求，模式为：并入&移交
type DeptToReq struct {
	Id     uint64 `json:"id" binding:"required" example:"1"`     // 当前部门数据ID
	ToId   uint64 `json:"toId" binding:"required" example:"1"`   // 目标部门数据ID
	ToType string `json:"toType" binding:"required" example:"0"` // 迁移模式，枚举：0_并入（子部门形式） 1_移交（部门保留，成员和子部门移交给新部门）
}
