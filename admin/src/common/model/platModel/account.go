package platModel

import (
	"fmt"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"siteol.com/smart/src/common/utils"
	"siteol.com/smart/src/common/utils/security"
	"strconv"
	"time"
)

// AccountDoReq 登陆账号 通用请求，创建&编辑可复用的字段
type AccountDoReq struct {
	Name           string   `json:"name" binding:"max=16" example:"demo"`                                // 姓名
	DeptId         string   `json:"deptId" binding:"required,number,gt=0" example:"0"`                   // 部门ID
	IsLeader       string   `json:"isLeader" binding:"oneof='0' '1'" example:"0"`                        // 部门职位，枚举：0_部门员工 1_部门领导
	PermissionType string   `json:"permissionType" binding:"required,oneof='0' '1' '2' '3'" example:"0"` // 权限类型，枚举：0_继承部门 1_本部门 2_本人 3_全局
	RoleIds        []uint64 `json:"roleIds" binding:"unique" example:"1,2,3"`                            // 关联的角色
	Status         string   `json:"status"  binding:"oneof='0' '1'" example:"0"`                         // 状态，枚举：0_正常 1_锁定 2_封存
}

// AccountAddReq 登陆账号 创建请求，酌情从通用中摘出部分字段
type AccountAddReq struct {
	Account string `json:"account" binding:"required,max=16" example:"demo"` // 登陆账号
	AccountDoReq
}

// AccountEditReq 登陆账号 编辑请求，酌情从通用中摘出部分字段
type AccountEditReq struct {
	Id uint64 `json:"id" binding:"required" example:"1"` // 数据ID
	AccountDoReq
}

// ToDbReq 登陆账号 创建转数据库
func (r *AccountAddReq) ToDbReq() (*platDB.Account, error) {
	now := time.Now()
	saltKey := utils.SaltKey()
	// 初始密码
	pwdC, err := security.AESEncrypt("123456", saltKey)
	if err != nil {
		return nil, err
	}
	deptId, _ := strconv.ParseUint(r.DeptId, 0, 64)
	return &platDB.Account{
		Id:             0,
		Account:        r.Account,
		Encryption:     pwdC,
		SaltKey:        saltKey,
		Name:           r.Name,
		DeptId:         deptId,
		IsLeader:       r.IsLeader,
		PermissionType: r.PermissionType,
		PwdExpTime:     &now, // 创建即超时
		LastLoginTime:  nil,
		Mark:           constant.StatusOpen,
		Status:         r.Status,
		CreateAt:       &now,
		UpdateAt:       &now,
	}, nil
}

// ToDbReq 登陆账号 更新转数据库
func (r *AccountEditReq) ToDbReq(d *platDB.Account) {
	now := time.Now()
	deptId, _ := strconv.ParseUint(r.DeptId, 0, 64)
	d.Name = r.Name
	d.DeptId = deptId
	d.IsLeader = r.IsLeader
	d.PermissionType = r.PermissionType
	d.Status = r.Status
	d.UpdateAt = &now
}

// AccountGetRes 登陆账号 详情响应
type AccountGetRes struct {
	Id             uint64   `json:"id" example:"1"`              // 数据ID
	Account        string   `json:"account" example:"demo"`      // 登陆账号
	Name           string   `json:"name" example:"demo"`         // 姓名
	DeptId         string   `json:"deptId" example:"0"`          // 部门ID
	DeptName       string   `json:"deptName" example:"DeptName"` // 部门名称
	IsLeader       string   `json:"isLeader" example:"0"`        // 部门职位，枚举：0_部门员工 1_部门领导
	PermissionType string   `json:"permissionType" example:"0"`  // 权限类型，枚举：0_继承部门 1_本部门 2_本人 3_全局
	PwdExpTime     string   `json:"pwdExpTime"`                  // 密码过期时间，创建即过期
	LastLoginTime  string   `json:"lastLoginTime"`               // 最后登陆时间，为空表示创建
	Status         string   `json:"status"`                      // 状态，枚举：0_正常 1_锁定 2_封存
	RoleIds        []uint64 `json:"roleIds"`                     // 关联的角色
	RoleNames      []string `json:"roleNames"`                   // 关联的角色名
}

// AccountPageReq 登陆账号 分页请求，根据实际业务替换分页条件字段
type AccountPageReq struct {
	Account string `json:"account" example:"demo"` // 登陆账号
	Name    string `json:"name" example:"demo"`    // 姓名
	DeptId  string `json:"deptId" example:"0"`     // 部门ID
	baseModel.PageReq
}

// AccountPageRes 登陆账号 分页响应，酌情从详情摘出部分字段
type AccountPageRes struct {
	AccountGetRes
}

// ToAccountGetRes 登陆账号 数据库转为详情响应
func ToAccountGetRes(r *platDB.Account, roleIds []uint64, roleMap, deptMap map[uint64]string) *AccountGetRes {
	roleNames := make([]string, len(roleIds))
	for i, item := range roleIds {
		roleNames[i] = roleMap[item]
	}
	return &AccountGetRes{
		Id:             r.Id,
		Account:        r.Account,
		Name:           r.Name,
		DeptId:         fmt.Sprintf("%d", r.DeptId),
		DeptName:       deptMap[r.DeptId],
		IsLeader:       r.IsLeader,
		PermissionType: r.PermissionType,
		RoleIds:        roleIds,
		RoleNames:      roleNames,
		Status:         r.Status,
	}
}

// ToAccountPageRes 登陆账号 数据库转分页响应
func ToAccountPageRes(list []*platDB.Account, roleIdList [][]uint64, roleMap, deptMap map[uint64]string) []*AccountPageRes {
	res := make([]*AccountPageRes, len(list))
	for i, r := range list {
		res[i] = &AccountPageRes{
			AccountGetRes: *ToAccountGetRes(r, roleIdList[i], roleMap, deptMap),
		}
	}
	return res
}
