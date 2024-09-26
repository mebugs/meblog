package platDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
	"time"
)

// LoginRecord 登陆记录
type LoginRecord struct {
	Id        uint64     `json:"id"`        // 默认数据ID
	AccountId uint64     `json:"accountId"` // 账号ID
	LoginType uint8      `json:"loginType"` // 登陆类型 1平台账号登录
	LoginTime *time.Time `json:"loginTime"` // 登陆时间
	Token     string     `json:"token"`     // 登陆Token
	Mark      string     `json:"mark"`      // 变更标识 0登陆成功 1主动登出 2被动登出
	Status    string     `json:"status"`    // 状态 0正常 1锁定 2封存
	CreateAt  *time.Time `json:"createAt"`  // 创建时间
	UpdateAt  *time.Time `json:"updateAt"`  // 更新时间
}

// LoginRecordTable 登陆记录泛型造器
var LoginRecordTable actuator.Table[LoginRecord]

// DataBase 实现指定数据库
func (t LoginRecord) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t LoginRecord) TableName() string {
	return "login_record"
}

// GetOutRangeRecords 获取限制外的登陆记录
func (t LoginRecord) GetOutRangeRecords(accountId uint64, limit uint64) (res []*LoginRecord, err error) {
	db := t.DataBase().Raw("SELECT * FROM `login_record` WHERE mark = 0 AND account_id = ? ORDER BY login_time DESC LIMIT ?,9999", accountId, limit).Scan(&res)
	err = db.Error
	return
}

// GetLoginRecordByToken 获取处于登陆状态的数据
func (t LoginRecord) GetLoginRecordByToken(token string) (res []*LoginRecord, err error) {
	db := t.DataBase().Raw("SELECT * FROM `login_record` WHERE token = ? AND mark = 0", token).Scan(&res)
	err = db.Error
	return
}
