package cacheModel

import (
	"encoding/json"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/mysql/platDB"
	"siteol.com/smart/src/common/redis"
)

// SyncRoleCache 同步全部角色
func SyncRoleCache(traceID string) (err error) {
	allList, err := platDB.RoleTable.GetAll()
	if err != nil {
		log.ErrorTF(traceID, "SyncRoleCache GetRole Fail . Err Is : %v", err)
		return
	}
	if len(allList) == 0 {
		log.WarnTF(traceID, "SyncRoleCache GetRole Empty")
		return
	}
	// 组装缓存对象
	resCacheMap := make(map[uint64]string, len(allList))
	for _, res := range allList {
		resCacheMap[res.Id] = res.Name
	}
	// 写入缓存 无超期
	err = redis.Set(constant.CacheRoles, resCacheMap, 0)
	if err != nil {
		log.InfoTF(traceID, "SyncRoleCache Fail . Err Is : %v", err)
	}
	return
}

// GetRoleCache 获取角色缓存
func GetRoleCache(traceID string) (res map[uint64]string, err error) {
	str, err := redis.Get(constant.CacheRoles)
	if err == nil {
		resCacheMap := make(map[uint64]string, 0)
		err = json.Unmarshal([]byte(str), &resCacheMap)
		if err == nil {
			res = resCacheMap
		} else {
			log.ErrorTF(traceID, "UnmarshalRoleCache Fail . Err Is : %v", err)
		}
	} else {
		log.ErrorTF(traceID, "GetRoleCache Fail . Err Is : %v", err)
	}
	return
}
