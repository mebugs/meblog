package blogDB

import (
	"siteol.com/smart/src/common/mysql/platDB"
	"time"

	"gorm.io/gorm"
	"siteol.com/smart/src/common/log"
)

var blogDb *gorm.DB

// InitBlogDb 初始化博客数据库
func InitBlogDb(traceId string) {
	blogDb = platDB.GetPlatDB()
	log.InfoTF(traceId, "Init BlogDb Success . ")
}

// Common 平台通用信息体
type Common struct {
	Mark     string     `json:"mark"`     // 变更标识 0可变更/登陆成功 1禁止变更/登陆失败
	Status   string     `json:"status"`   // 状态 0正常 1锁定 2封存
	CreateAt *time.Time `json:"createAt"` // 创建时间
	UpdateAt *time.Time `json:"updateAt"` // 更新时间
}
