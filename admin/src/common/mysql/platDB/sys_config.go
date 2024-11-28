package platDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
)

// SysConfig 系统配置
type SysConfig struct {
	Id                uint64 `json:"id"`                // 数据ID
	LoginSwitch       string `json:"loginSwitch"`       // 并发限制开关，0限制 1不限制
	LoginNum          uint16 `json:"loginNum"`          // 最大登陆并发量，最小为1
	LoginFailSwitch   string `json:"loginFailSwitch"`   // 登陆失败限制开关，0限制 1不限制
	LoginFailUnit     string `json:"loginFailUnit"`     // 登陆失败限制 1秒 2分 3时 4天
	LoginFailNum      uint16 `json:"loginFailNum"`      // 登陆失败最大尝试次数
	LoginFailLockUnit string `json:"loginFailLockUnit"` // 登陆失败锁定 1秒 2分 3时 4天
	LoginFailLockNum  uint16 `json:"loginFailLockNum"`  // 登陆失败锁定时长
	LoginFailTryNum   uint16 `json:"loginFailTryNum"`   // 登陆失败尝试次数
	LogoutSwitch      string `json:"logoutSwitch"`      // 登陆过期开关，0限制 1不限制
	LogoutUnit        string `json:"logoutUnit"`        // 登陆过期单位，0永不过期 1秒 2分 3时 4天
	LogoutNum         uint16 `json:"logoutNum"`         // 登陆过期长度数量
	FileFullPath      string `json:"fileFullPath"`      // 文件存放根目录
}

// SysConfigTable 系统配置泛型造器
var SysConfigTable actuator.Table[SysConfig]

// DataBase 实现指定数据库
func (t SysConfig) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t SysConfig) TableName() string {
	return "sys_config"
}
