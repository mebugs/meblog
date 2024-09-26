package cacheModel

import (
	"encoding/json"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/mysql/platDB"
	"siteol.com/smart/src/common/redis"
)

// SyncPermissionCache 同步全部权限
func SyncPermissionCache(traceID string) (err error) {
	allList, err := platDB.PermissionTable.GetAll()
	if err != nil {
		log.ErrorTF(traceID, "SyncPermissionCache GetPermission Fail . Err Is : %v", err)
		return
	}
	if len(allList) == 0 {
		log.WarnTF(traceID, "SyncPermissionCache GetPermission Empty")
		return
	}
	// 组装缓存对象
	// 默认权限ID，所有角色/账号，默认赋予
	normalPermIds := make([]uint64, 0)
	resCacheMap := make(map[uint64]string, len(allList))
	for _, res := range allList {
		resCacheMap[res.Id] = res.Alias
		if res.Static == constant.StatusOpen {
			// 默认权限
			normalPermIds = append(normalPermIds, res.Id)
		}
	}
	// 写入缓存 无超期
	err = redis.Set(constant.CachePermissions, resCacheMap, 0)
	if err != nil {
		log.InfoTF(traceID, "SyncPermissionCache Fail . Err Is : %v", err)
	}
	err = redis.Set(constant.CachePermissionsNormal, normalPermIds, 0)
	if err != nil {
		log.InfoTF(traceID, "SyncPermissionCache Fail . Err Is : %v", err)
	}
	return
}

// GetPermissionCache 获取权限缓存
func GetPermissionCache(traceID string) (res map[uint64]string, err error) {
	str, err := redis.Get(constant.CachePermissions)
	if err == nil {
		resCacheMap := make(map[uint64]string, 0)
		err = json.Unmarshal([]byte(str), &resCacheMap)
		if err == nil {
			res = resCacheMap
		} else {
			log.ErrorTF(traceID, "UnmarshalPermissionCache Fail . Err Is : %v", err)
		}
	} else {
		log.ErrorTF(traceID, "GetPermissionCache Fail . Err Is : %v", err)
	}
	return
}

// GetPermissionNormalIdsCache 获取默认权限缓存
func GetPermissionNormalIdsCache(traceID string) (res []uint64, err error) {
	str, err := redis.Get(constant.CachePermissionsNormal)
	if err == nil {
		resCache := make([]uint64, 0)
		err = json.Unmarshal([]byte(str), &resCache)
		if err == nil {
			res = resCache
		} else {
			log.ErrorTF(traceID, "UnmarshalPermissionNormalCache Fail . Err Is : %v", err)
		}
	} else {
		log.ErrorTF(traceID, "GetPermissionNormalCache Fail . Err Is : %v", err)
	}
	return
}
