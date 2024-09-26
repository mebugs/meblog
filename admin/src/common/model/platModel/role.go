package platModel

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"time"
)

// RoleDoReq 角色通用请求
type RoleDoReq struct {
	Name              string   `json:"name" binding:"required,max=16" example:"Admin"`     // 角色名称
	Remark            string   `json:"remark" binding:"max=64" example:"This is Admin"`    // 角色备注
	PermissionIds     []uint64 `json:"permissionIds" binding:"unique" example:"1,2,3"`     // 权限集
	HalfPermissionIds []uint64 `json:"halfPermissionIds" binding:"unique" example:"1,2,3"` // 权限集（父级半选）
}

// RoleAddReq 角色创建请求
type RoleAddReq struct {
	RoleDoReq
}

// ToDbReq 角色创建对象转字典对象
func (r *RoleAddReq) ToDbReq() *platDB.Role {
	now := time.Now()
	return &platDB.Role{
		Id:     0,
		Name:   r.Name,
		Remark: r.Remark,
		Common: platDB.Common{
			Mark:     constant.StatusOpen,
			Status:   constant.StatusOpen,
			CreateAt: &now,
			UpdateAt: &now,
		},
	}
}

// RoleEditReq 角色编辑请求
type RoleEditReq struct {
	Id uint64 `json:"id" binding:"required" example:"1"` // 数据ID
	RoleDoReq
}

// ToDbReq 角色更新对象转字典对象
func (r *RoleEditReq) ToDbReq(d *platDB.Role) {
	now := time.Now()
	d.Name = r.Name
	d.Remark = r.Remark
	d.UpdateAt = &now
}

// RoleGetRes 角色分页响应
type RoleGetRes struct {
	Id                uint64   `json:"id" example:"1"`                    // 数据ID
	Name              string   `json:"name" example:"Admin"`              // 角色名称
	Remark            string   `json:"remark" example:"This is Admin"`    // 角色备注
	PermissionIds     []uint64 `json:"permissionIds" example:"1,2,3"`     // 权限集
	HalfPermissionIds []uint64 `json:"halfPermissionIds" example:"1,2,3"` // 权限集（父级半选）

}

// ToRoleGetRes 角色转为查询对象
func ToRoleGetRes(r *platDB.Role, permissions []*platDB.RolePermission) *RoleGetRes {
	ids := make([]uint64, 0)
	halfIds := make([]uint64, 0)
	if len(permissions) > 0 {
		for _, item := range permissions {
			// 选中
			if item.CheckType == constant.StatusOpen {
				ids = append(ids, item.PermissionId)
			} else {
				halfIds = append(halfIds, item.PermissionId)
			}
		}
	}
	return &RoleGetRes{
		Id:                r.Id,
		Name:              r.Name,
		Remark:            r.Remark,
		PermissionIds:     ids,
		HalfPermissionIds: halfIds,
	}
}

// RolePageReq 角色分页请求
type RolePageReq struct {
	Name string `json:"name" example:"Admin"` // 角色名称，支持模糊搜索
	baseModel.PageReq
}

// RolePageRes 角色分页响应
type RolePageRes struct {
	RoleGetRes
	Mark string `json:"mark" example:"1"` // 变更标识 0可变更 1禁止变更
}

// ToRolePageRes 角色转为分页对象
func ToRolePageRes(list []*platDB.Role) []*RolePageRes {
	res := make([]*RolePageRes, len(list))
	for i, r := range list {
		res[i] = &RolePageRes{
			RoleGetRes: *ToRoleGetRes(r, nil),
			Mark:       r.Mark,
		}
	}
	return res
}
