package platModel

import "siteol.com/smart/src/common/mysql/platDB"

// SysConfigGetRes 系统配置查询结果
type SysConfigGetRes struct {
	LoginSwitch       string `json:"loginSwitch" example:"0"`       // 并发限制开关，0限制 1不限制
	LoginNum          uint16 `json:"loginNum" example:"1"`          // 最大登陆并发量，最小为1
	LoginFailSwitch   string `json:"loginFailSwitch" example:"0"`   // 登陆失败限制开关，0限制 1不限制
	LoginFailUnit     string `json:"loginFailUnit" example:"1"`     // 登陆失败限制 1秒 2分 3时 4天
	LoginFailNum      uint16 `json:"loginFailNum" example:"1"`      // 登陆失败最大尝试次数，最小为1
	LoginFailLockUnit string `json:"loginFailLockUnit" example:"1"` // 登陆失败锁定 1秒 2分 3时 4天
	LoginFailLockNum  uint16 `json:"loginFailLockNum" example:"1"`  // 登陆失败锁定数量，最小为1
	LoginFailTryNum   uint16 `json:"loginFailTryNum" example:"1"`   // 登陆失败尝试次数
	LogoutSwitch      string `json:"logoutSwitch" example:"0"`      // 登陆过期开关，0限制 1不限制
	LogoutUnit        string `json:"logoutUnit" example:"1"`        // 登陆过期单位，1秒 2分 3时 4天
	LogoutNum         uint16 `json:"logoutNum" example:"1"`         // 登陆过期长度数量，最小为1
}

// ToSysConfigGetRes 系統配置转为查询对象
func ToSysConfigGetRes(r *platDB.SysConfig) *SysConfigGetRes {
	return &SysConfigGetRes{
		LoginSwitch:       r.LoginSwitch,
		LoginNum:          r.LoginNum,
		LoginFailSwitch:   r.LoginFailSwitch,
		LoginFailUnit:     r.LoginFailUnit,
		LoginFailNum:      r.LoginFailNum,
		LoginFailLockUnit: r.LoginFailLockUnit,
		LoginFailLockNum:  r.LoginFailLockNum,
		LoginFailTryNum:   r.LoginFailTryNum,
		LogoutSwitch:      r.LogoutSwitch,
		LogoutUnit:        r.LogoutUnit,
		LogoutNum:         r.LogoutNum,
	}
}

// SysConfigEditReq 系统配置编辑请求
type SysConfigEditReq struct {
	LoginSwitch       string `json:"loginSwitch" binding:"required,oneof='0' '1'" example:"0"`     // 并发限制开关，0限制 1不限制
	LoginNum          uint16 `json:"loginNum" binding:"number" example:"1"`                        // 最大登陆并发量，最小为1
	LoginFailSwitch   string `json:"loginFailSwitch" binding:"required,oneof='0' '1'" example:"0"` // 登陆失败限制开关，0限制 1不限制
	LoginFailUnit     string `json:"loginFailUnit" binding:"number" example:"1"`                   // 登陆失败限制 1秒 2分 3时 4天
	LoginFailNum      uint16 `json:"loginFailNum" binding:"number" example:"1"`                    // 登陆失败限制数，最小为1
	LoginFailLockUnit string `json:"loginFailLockUnit" binding:"number" example:"1"`               // 登陆失败锁定 1秒 2分 3时 4天
	LoginFailLockNum  uint16 `json:"loginFailLockNum" binding:"number" example:"1"`                // 登陆失败锁定数量，最小为1
	LoginFailTryNum   uint16 `json:"loginFailTryNum" binding:"number" example:"1"`                 // 登陆失败尝试次数
	LogoutSwitch      string `json:"logoutSwitch" binding:"required,oneof='0' '1'" example:"0"`    // 登陆过期开关，0限制 1不限制
	LogoutUnit        string `json:"logoutUnit" binding:"number" example:"1"`                      // 登陆过期单位，1秒 2分 3时 4天
	LogoutNum         uint16 `json:"logoutNum" binding:"number" example:"1"`                       // 登陆过期长度数量，最小为1
}

// ToDbReq 字典更新对象转字典对象
func (r *SysConfigEditReq) ToDbReq(d *platDB.SysConfig) {
	d.LoginSwitch = r.LoginSwitch
	d.LoginNum = r.LoginNum
	d.LoginFailSwitch = r.LoginFailSwitch
	d.LoginFailUnit = r.LoginFailUnit
	d.LoginFailNum = r.LoginFailNum
	d.LoginFailLockUnit = r.LoginFailLockUnit
	d.LoginFailLockNum = r.LoginFailLockNum
	d.LoginFailTryNum = r.LoginFailTryNum
	d.LogoutSwitch = r.LogoutSwitch
	d.LogoutUnit = r.LogoutUnit
	d.LogoutNum = r.LogoutNum
}
