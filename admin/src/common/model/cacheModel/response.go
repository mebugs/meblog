package cacheModel

import (
	"encoding/json"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/mysql/platDB"
	"siteol.com/smart/src/common/redis"
)

// SyncResponseCache 同步响应码配置
func SyncResponseCache(traceID string) (err error) {
	allList, err := platDB.ResponseTable.GetAll()
	if err != nil {
		log.ErrorTF(traceID, "SyncResponseCache GetResponse Fail . Err Is : %v", err)
		return
	}
	if len(allList) == 0 {
		log.WarnTF(traceID, "SyncResponseCache GetResponse Empty")
		return
	}
	// 组装缓存对象
	resCodeCacheMap := make(map[string]map[string]string, 0)
	for _, res := range allList {
		if res.Status != constant.StatusOpen {
			continue
		}
		// 遍历支持的语言并写入Map
		langMap := make(map[string]string, len(constant.TransLangSupport))
		for _, lang := range constant.TransLangSupport {
			switch lang {
			case "en-US":
				langMap[lang] = res.EnUs
			case "zh-CN":
				langMap[lang] = res.ZhCn
			}
		}
		resCodeCacheMap[res.Code] = langMap
	}
	// 写入缓存 无超期
	err = redis.Set(constant.CacheResTrans, resCodeCacheMap, 0)
	if err != nil {
		log.InfoTF(traceID, "SyncResponseCache Fail . Err Is : %v", err)
	}
	return
}

// GetResponseCache 获取响应码配置
func GetResponseCache(traceID string) (transMap map[string]map[string]string, err error) {
	// 读取缓存
	tranStr, err := redis.Get(constant.CacheResTrans)
	if err != nil {
		log.WarnTF(traceID, "GetResponseCacheLangCacheMap Fail . Err Is : %v", err)
		// 出错不翻译
		return

	}
	transMap = make(map[string]map[string]string, 0)
	err = json.Unmarshal([]byte(tranStr), &transMap)
	if err != nil {
		log.ErrorTF(traceID, "Unmarshal ResponseCacheLangCacheMap Fail . Err Is : %v", err)
		// 出错不翻译
	}
	return
}
