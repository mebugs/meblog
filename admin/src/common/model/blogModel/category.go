package blogModel

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/blogDB"
	"time"
)

// CategoryDoReq 文章分类 通用请求，创建&编辑可复用的字段
type CategoryDoReq struct {
	Title   string `json:"title" binding:"max=32" example:"demo"`   // 名称
	Url     string `json:"url" binding:"max=32" example:"demo"`     // 分类地址
	Summary string `json:"summary" binding:"max=64" example:"demo"` // 简介
}

// CategoryAddReq 文章分类 创建请求，酌情从通用中摘出部分字段
type CategoryAddReq struct {
	Source []string `json:"source"  binding:"required" example:"0"` // 主图资源列表
	CategoryDoReq
}

// CategoryEditReq 文章分类 编辑请求，酌情从通用中摘出部分字段
type CategoryEditReq struct {
	Id     uint64   `json:"id" binding:"required" example:"1"` // 数据ID
	Source []string `json:"source" example:"0"`                // 主图资源列表
	CategoryDoReq
}

// ToDbReq 文章分类 创建转数据库
func (r *CategoryAddReq) ToDbReq() *blogDB.Category {
	now := time.Now()
	return &blogDB.Category{
		Id:       0,
		Title:    r.Title,
		Url:      r.Url,
		Summary:  r.Summary,
		Status:   constant.StatusOpen,
		CreateAt: &now,
		UpdateAt: &now,
	}
}

// ToDbReq 文章分类 更新转数据库
func (r *CategoryEditReq) ToDbReq(d *blogDB.Category) {
	now := time.Now()
	d.Title = r.Title
	d.Url = r.Url
	d.Summary = r.Summary
	d.UpdateAt = &now
}

// CategoryGetRes 文章分类 详情响应
type CategoryGetRes struct {
	Id         uint64 `json:"id" example:"1"`                // 数据ID
	Title      string `json:"title" example:"demo"`          // 名称
	Url        string `json:"url" example:"demo"`            // 分类地址
	Summary    string `json:"summary" example:"demo"`        // 简介
	SourceShow string `json:"sourceShow" example:"/xxx.jpg"` // 资源图片地址
	Status     string `json:"status" example:"0"`            // 状态，枚举：0_正常 1_锁定 2_封存
}

// CategoryPageReq 文章分类 分页请求，根据实际业务替换分页条件字段
type CategoryPageReq struct {
	Id uint64 `json:"id" example:"1"` // 数据ID
	baseModel.PageReq
}

// CategoryPageRes 文章分类 分页响应，酌情从详情摘出部分字段
type CategoryPageRes struct {
	CategoryGetRes
}

// ToCategoryGetRes 文章分类 数据库转为详情响应
func ToCategoryGetRes(r *blogDB.Category, sourceShow string) *CategoryGetRes {
	return &CategoryGetRes{
		Id:         r.Id,
		Title:      r.Title,
		Url:        r.Url,
		Summary:    r.Summary,
		SourceShow: sourceShow,
		Status:     r.Status,
	}
}

// ToCategoryPageRes 文章分类 数据库转分页响应
func ToCategoryPageRes(list []*blogDB.Category, sourceMap map[uint64]string) []*CategoryPageRes {
	res := make([]*CategoryPageRes, len(list))
	for i, r := range list {
		res[i] = &CategoryPageRes{
			CategoryGetRes: *ToCategoryGetRes(r, sourceMap[r.SourceId]),
		}
	}
	return res
}

// CategoryToReq 分组迁移至另一分组
type CategoryToReq struct {
	Id   uint64 `json:"id" binding:"required" example:"1"`              // 当前分组数据ID
	ToId uint64 `json:"toId" binding:"required,nefield=Id" example:"1"` // 目标分组数据ID
}
