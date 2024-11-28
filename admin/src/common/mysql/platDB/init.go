package platDB

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/config"
)

var platDb *gorm.DB

// InitPlatFromDb 初始化平台数据库
func InitPlatFromDb(traceId string) {
	// 采用默认配置打开数据可（默认禁用事务）
	db, err := gorm.Open(mysql.Open(config.JsonConfig.MySQL.Plat), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.FatalTF(fmt.Sprintf("%s%s", config.SysNode, "INIT"), "Open PlatDb Fail . Err Is : %s", err)
		return
	}
	platDb = db
	log.InfoTF(traceId, "Init PlatDb Success . ")
}

// GetPlatDB 博客和平台共用了数据库，所以直接引用
func GetPlatDB() *gorm.DB {
	return platDb
}

// Common 平台通用信息体
type Common struct {
	Mark     string     `json:"mark"`     // 变更标识 0可变更/登陆成功 1禁止变更/登陆失败
	Status   string     `json:"status"`   // 状态 0正常 1锁定 2封存
	CreateAt *time.Time `json:"createAt"` // 创建时间
	UpdateAt *time.Time `json:"updateAt"` // 更新时间
}
