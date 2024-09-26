package cacheModel

import (
	"encoding/json"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/mysql/platDB"
	"siteol.com/smart/src/common/redis"
	"siteol.com/smart/src/common/utils"
)

var CacheRouterNormal = &CacheRouter{}

// CacheRouter 缓存路由对象
type CacheRouter struct {
	Id        uint64   `json:"id"`        // 数据ID，为0表示路由未配置
	NeedAuth  bool     `json:"needAuth"`  // 是否需要授权
	LogInDb   bool     `json:"logInDb"`   // 日志入库
	ReqPrint  bool     `json:"reqPrint"`  // 请求日志打印
	ReqSecure []string `json:"reqSecure"` // 请求日志脱敏数组字段
	ResPrint  bool     `json:"resPrint"`  // 响应日志打印
	ResSecure []string `json:"resSecure"` // 响应日志脱敏数组字段
}

// SyncRouterCache 同步路由配置
func SyncRouterCache(traceID string) (err error) {
	allList, err := platDB.RouterTable.GetAll()
	if err != nil {
		log.ErrorTF(traceID, "SyncRouterCache GetRouter Fail . Err Is : %v", err)
		return
	}
	if len(allList) == 0 {
		log.WarnTF(traceID, "SyncRouterCache GetRouter Empty")
		return
	}
	// 组装缓存对象
	resCacheMap := make(map[string]*CacheRouter, 0)
	urlMap := make(map[uint64]string, 0)
	for _, res := range allList {
		// 路由是物理删除
		cache := &CacheRouter{
			Id:        res.Id,
			NeedAuth:  utils.StatusBool(res.Type),
			LogInDb:   utils.StatusBool(res.LogInDb),
			ReqPrint:  utils.StatusBool(res.ReqLogPrint),
			ReqSecure: utils.ArrayStr(res.ReqLogSecure),
			ResPrint:  utils.StatusBool(res.ResLogPrint),
			ResSecure: utils.ArrayStr(res.ResLogSecure),
		}
		resCacheMap[res.Url] = cache
		urlMap[res.Id] = res.Url
	}
	// 写入缓存 无超期
	err = redis.Set(constant.CacheRouters, resCacheMap, 0)
	if err != nil {
		log.InfoTF(traceID, "SyncRouterCache Fail . Err Is : %v", err)
	}
	err = redis.Set(constant.CacheRouterUrls, urlMap, 0)
	if err != nil {
		log.InfoTF(traceID, "SyncRouterUrlCache Fail . Err Is : %v", err)
	}
	return
}

// GetRouterCache 获取路由配置
func GetRouterCache(traceID string) (cache map[string]*CacheRouter, err error) {
	str, err := redis.Get(constant.CacheRouters)
	if err != nil {
		log.ErrorTF(traceID, "GetRouterCache Fail . Err Is : %v", err)
		return
	}
	cache = make(map[string]*CacheRouter, 0)
	err = json.Unmarshal([]byte(str), &cache)
	if err != nil {
		log.ErrorTF(traceID, "Unmarshal RouterCache Fail . Err Is : %v", err)
	}
	return
}

// GetRouterUrlsCache 获取路由地址配置
func GetRouterUrlsCache(traceID string) (cache map[uint64]string, err error) {
	str, err := redis.Get(constant.CacheRouterUrls)
	if err != nil {
		log.ErrorTF(traceID, "GetRouterUrlCache Fail . Err Is : %v", err)
		return
	}
	cache = make(map[uint64]string, 0)
	err = json.Unmarshal([]byte(str), &cache)
	if err != nil {
		log.ErrorTF(traceID, "Unmarshal RouterUrlCache Fail . Err Is : %v", err)
	}
	return
}
