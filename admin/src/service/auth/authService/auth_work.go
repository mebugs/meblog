package authService

import (
	"fmt"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/cacheModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"siteol.com/smart/src/common/redis"
	"strconv"
	"time"
)

// checkLoginFailRule 检查登陆风控规则
func checkLoginFailRule(acc string, sysConf *cacheModel.CacheSysConfig) (get bool) {
	// 未开启风控
	if !sysConf.LoginFailSwitch {
		return
	}
	// 查询统计指数
	loginLockStr, err := redis.Get("LoginLock::" + acc)
	if err != nil {
		return
	}
	// 正被锁定
	if loginLockStr == "1" {
		get = true
	}
	return
}

// syncLoginFailRule 刷新登陆风控
func syncLoginFailRule(traceID, acc string, sysConf *cacheModel.CacheSysConfig) {
	// 未开启风控
	if !sysConf.LoginFailSwitch {
		return
	}
	// 查询统计指数
	loginTryStr, err := redis.Get("LoginTry::" + acc)
	if err != nil {
		loginTryStr = "0"
	}
	// 指数+1
	loginTryNum, err := strconv.ParseUint(loginTryStr, 10, 64)
	if err != nil {
		return
	}
	loginTryNum++
	// 刷新登陆失败的同级结果和周期
	_ = redis.Set("LoginTry::"+acc, loginTryNum, sysConf.LoginFailLimit)
	// 指数判定
	if loginTryNum > sysConf.LoginFailTryNum {
		// 增加风控拦截
		_ = redis.Set("LoginLock::"+acc, "1", sysConf.LoginFailLockLimit)
		log.WarnTF(traceID, "Account %s Has Been Add LoginFailRule", acc)
	}
	return
}

// syncLoginRecord 同步登陆记录和踢出
func syncLoginRecord(traceID, token string, accountId uint64, sysConf *cacheModel.CacheSysConfig) {
	now := time.Now()
	record := &platDB.LoginRecord{
		AccountId: accountId,
		LoginType: 1, // 平台登陆
		LoginTime: &now,
		Token:     token,
		Mark:      constant.StatusOpen, // 登陆成功
		Status:    constant.StatusOpen,
		CreateAt:  &now,
		UpdateAt:  &now,
	}
	// 插入登陆记录
	err := platDB.LoginRecordTable.InsertOne(record)
	if err != nil {
		log.WarnTF(traceID, "Insert LoginRecord Fail . Err Is : %v", err)
	}
	// 检证多端踢出，未启用或配置为0属于不检证
	if !sysConf.LoginSwitch || sysConf.LoginNum < 1 {
		return
	}
	// 检索出多余的历史登陆数据
	records, err := platDB.LoginRecordTable.Executor().GetOutRangeRecords(accountId, sysConf.LoginNum)
	if err != nil {
		log.WarnTF(traceID, "GetOutRangeLoginRecords Fail . Err Is : %v", err)
		return
	}
	// 遍历数据，踢出缓存，批量更新
	ids := make([]uint64, len(records))
	for i, item := range records {
		_ = redis.Del(fmt.Sprintf(constant.CacheAuth, item.Token))
		ids[i] = item.Id
	}
	// 批量更新
	err = platDB.LoginRecordTable.UpdateByIds(ids, map[string]any{
		"mark":      constant.StatusClose, // 被动登出
		"update_at": &now,
	})
	if err != nil {
		log.WarnTF(traceID, "UpdateOutRangeLoginRecords Fail . Err Is : %v", err)
	}
	return
}
