package cacheModel

import (
	"encoding/json"
	"fmt"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/redis"
)

// CacheAuth 缓存授权对象
type CacheAuth struct {
	AccountId          uint64   `json:"accountId"`          // 账号ID，当数据权限为个人时，账号ID将进行数据过滤
	Name               string   `json:"name"`               // 账号ID
	RoleIds            []uint64 `json:"roleIds"`            // 角色ID列表
	RoleNames          []string `json:"roleNames"`          // 角色名列表
	PermissionIds      []uint64 `json:"permissionIds"`      // 权限ID列表
	PermissionKeys     []string `json:"permissionKeys"`     // 权限KEY列表（前端使用）
	Routers            []string `json:"routers"`            // 服务端路由列表
	DataPermissionType string   `json:"dataPermissionType"` // 权限类型，枚举：0_本部门与子部门 1_本部门 2_个人 3_全局
	DeptIds            []uint64 `json:"deptIds"`            // 响应日志脱敏数组字段
	NeedResetPwd       bool     `json:"NeedResetPwd"`       // 需要进行密码重置
}

// GetAuthCache 获取授权用户缓存
func GetAuthCache(traceID, token string) (authCache *CacheAuth, err error) {
	// 获取授权用户缓存
	authStr, err := redis.Get(fmt.Sprintf(constant.CacheAuth, token))
	if err != nil {
		log.ErrorTF(traceID, "Get AuthCache Fail . Err Is : %v", err)
		return
	}
	// 解析成对象
	authCache = &CacheAuth{}
	err = json.Unmarshal([]byte(authStr), authCache)
	if err != nil {
		log.ErrorTF(traceID, "GetAuthCache JSONUnmarshal Fail. Err Is : %v", err)
	}
	return
}

// RefreshAuthCache 根据配置刷新授权时间
func RefreshAuthCache(traceID, token string) (err error) {
	sysConf, err := GetSysConfigCache(traceID)
	if err != nil {
		return
	}
	// 根据配置，写入Redis
	outTime := uint64(0)
	if sysConf.LogoutSwitch {
		outTime = sysConf.LogoutLimit
	}
	// 写入缓存，刷新过期时间
	err = redis.SetTTL(fmt.Sprintf(constant.CacheAuth, token), outTime)
	if err != nil {
		log.ErrorTF(traceID, "RefreshAuthCache Fail. Err Is : %v", err)
	}
	return
}
