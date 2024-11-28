package cacheModel

import (
	"encoding/json"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/mysql/platDB"
	"siteol.com/smart/src/common/redis"
	"siteol.com/smart/src/common/utils"
)

// CacheSysConfig 缓存系统配置
type CacheSysConfig struct {
	Id                 uint64 `json:"id"`                 // 数据ID
	LoginSwitch        bool   `json:"loginSwitch"`        // 并发限制开关
	LoginNum           uint64 `json:"loginNum"`           // 最大登陆并发量，最小为1
	LoginFailSwitch    bool   `json:"loginFailSwitch"`    // 登陆失败限制开关
	LoginFailLimit     uint64 `json:"loginFailLimit"`     // 登陆失败限制，计数周期
	LoginFailLockLimit uint64 `json:"loginFailLockLimit"` // 登陆失败限制，锁定周期
	LoginFailTryNum    uint64 `json:"loginFailTryNum"`    // 登陆失败尝试次数
	LogoutSwitch       bool   `json:"logoutSwitch"`       // 登陆过期开关
	LogoutLimit        uint64 `json:"logoutLimit"`        // 登陆过期单位，刷新周期
	FileFullPath       string `json:"fileFullPath"`       // 文件存放根目录
}

// SyncSysConfigCache 同步系统配置
func SyncSysConfigCache(traceID string) (err error) {
	db, err := platDB.SysConfigTable.GetOneById(1)
	if err != nil {
		log.ErrorTF(traceID, "GetSysConfig Fail . Err Is : %v", err)
		return
	}
	// 转换成缓存对象
	dbCache := &CacheSysConfig{
		Id:                 db.Id,
		LoginSwitch:        utils.StatusBool(db.LoginSwitch),
		LoginNum:           uint64(db.LoginNum),
		LoginFailSwitch:    utils.StatusBool(db.LoginFailSwitch),
		LoginFailLimit:     utils.GetMilliseconds(db.LoginFailNum, db.LoginFailUnit),
		LoginFailLockLimit: utils.GetMilliseconds(db.LoginFailLockNum, db.LoginFailLockUnit),
		LoginFailTryNum:    uint64(db.LoginFailTryNum),
		LogoutSwitch:       utils.StatusBool(db.LogoutSwitch),
		LogoutLimit:        utils.GetMilliseconds(db.LogoutNum, db.LogoutUnit),
		FileFullPath:       db.FileFullPath,
	}
	err = redis.Set(constant.CacheSysConf, dbCache, 0)
	if err != nil {
		log.ErrorTF(traceID, "SetSysConfigCache Fail . Err Is : %v", err)
	}
	return
}

// GetSysConfigCache 获取系统配置
func GetSysConfigCache(traceID string) (sysConf *CacheSysConfig, err error) {
	sysConfStr, err := redis.Get(constant.CacheSysConf)
	if err != nil {
		log.ErrorTF(traceID, "GetSysConfigCache Fail . Err Is : %v", err)
		return
	}
	sysConf = &CacheSysConfig{}
	err = json.Unmarshal([]byte(sysConfStr), sysConf)
	if err != nil {
		log.ErrorTF(traceID, "GetSysConfigCache JSONUnmarshal Fail. Err Is : %v", err)
	}
	return
}
