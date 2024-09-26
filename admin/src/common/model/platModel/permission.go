package platModel

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"time"
)

// PermissionDoReq 权限通用请求
type PermissionDoReq struct {
	Name      string   `json:"name" binding:"required,max=32" example:"Center"`  // 权限名称，界面展示，建议与界面导航一致
	Alias     string   `json:"alias" binding:"required,max=32" example:"Center"` // 权限别名，英文，规范如下：sys，sysAccount sysAccountAdd
	RouterIds []uint64 `json:"routerIds" binding:"unique" example:"1,2,3"`       // 路由集，提交路由ID数组
}

// PermissionAddReq 权限创建请求
type PermissionAddReq struct {
	PermissionDoReq
	Level  string `json:"level" binding:"required,max=1" example:"1"`          // 权限等级 1分组（一级导航）2模块（页面）3功能（按钮）第四级路由不在本表中体现
	Pid    uint64 `json:"pid" binding:"required" example:"1"`                  // 父级ID，默认为1
	Static string `json:"static" binding:"required,oneof='0' '1'" example:"1"` // 默认启用权限，0 启用 1 不启，启用后，该权限默认被分配，不可去勾
}

// ToDbReq 权限创建对象转字典对象
func (r *PermissionAddReq) ToDbReq() *platDB.Permission {
	now := time.Now()
	return &platDB.Permission{
		Id:     0,
		Name:   r.Name,
		Alias:  r.Alias,
		Level:  r.Level,
		Pid:    r.Pid,
		Sort:   0,
		Static: r.Static,
		Common: platDB.Common{
			Mark:     constant.StatusOpen,
			Status:   constant.StatusOpen,
			CreateAt: &now,
			UpdateAt: &now,
		},
	}
}

// PermissionEditReq 权限编辑请求
type PermissionEditReq struct {
	Id uint64 `json:"id" binding:"required" example:"1"` // 数据ID
	PermissionDoReq
}

// ToDbReq 权限更新对象转字典对象
func (r *PermissionEditReq) ToDbReq(d *platDB.Permission) {
	now := time.Now()
	d.Name = r.Name
	d.Alias = r.Alias
	d.UpdateAt = &now
}

// PermissionGetRes 权限分页响应
type PermissionGetRes struct {
	Id        uint64           `json:"id" example:"1"`            // 数据ID
	Name      string           `json:"name" example:"Center"`     // 权限名称，界面展示，建议与界面导航一致
	Alias     string           `json:"alias" example:"Center"`    // 权限别名，英文，规范如下：sys，sysAccount sysAccountAdd
	Level     string           `json:"level" example:"1"`         // 权限等级 1分组（一级导航）2模块（页面）3功能（按钮）第四级路由不在本表中体现
	Pid       uint64           `json:"pid" example:"1"`           // 父级ID，默认为1
	Static    string           `json:"static" example:"1"`        // 默认启用权限，0 启用 1 不启，启用后，该权限默认被分配，不可去勾
	RouterIds []uint64         `json:"routerIds" example:"1,2,3"` // 路由集，提交路由ID数组
	Routers   []*RouterPageRes `json:"routers"`                   // 路由列表
}

// ToPermissionGetRes 权限转为查询对象
func ToPermissionGetRes(r *platDB.Permission, routerIds []uint64, routers []*RouterPageRes) *PermissionGetRes {
	return &PermissionGetRes{
		Id:        r.Id,
		Name:      r.Name,
		Alias:     r.Alias,
		Level:     r.Level,
		Pid:       r.Pid,
		Static:    r.Static,
		RouterIds: routerIds,
		Routers:   routers,
	}
}

// ToPermissionBroRes 权限转为查询对象
func ToPermissionBroRes(r platDB.PermissionArray) []*baseModel.SortRes {
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
