package blogDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
	"time"
)

// Post 文章或页面
type Post struct {
	Id         uint64     `json:"id"`         // 数据ID
	Title      string     `json:"title"`      // 标题
	Url        string     `json:"url"`        // 文章地址
	Summary    string     `json:"summary"`    // 摘要
	SourceId   uint64     `json:"sourceId"`   // 主图资源ID
	PostType   uint8      `json:"postType"`   // 文章类型，枚举：0_页面 1_文章
	CategoryId uint64     `json:"categoryId"` // 分组ID，文章可用
	Md         string     `json:"md"`         // MD源文件
	Html       string     `json:"html"`       // HTML源文件
	PushAt     *time.Time `json:"pushAt"`     // 发布时间，可以自动任务抓取，不配代表不自动发布
	Status     string     `json:"status"`     // 状态，枚举：0_正常 1_锁定 2_封存
	CreateAt   *time.Time `json:"createAt"`   // 创建时间
	UpdateAt   *time.Time `json:"updateAt"`   // 更新时间
}

// PostTable 文章或页面泛型造器
var PostTable actuator.Table[Post]

// DataBase 实现指定数据库
func (t Post) DataBase() *gorm.DB {
	return blogDb
}

// TableName 实现自定义表名
func (t Post) TableName() string {
	return "post"
}

// ToCategory 迁移到新分组
func (t Post) ToCategory(id, toId uint64) (err error) {
	now := time.Now()
	r := blogDb.Table(t.TableName()).Where("category_id = ?", id).Updates(map[string]any{
		"category_id": toId,
		"update_at":   &now,
	})
	err = r.Error
	return
}
