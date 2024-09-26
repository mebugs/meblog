package platDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
)

// Dict 字典
type Dict struct {
	Id       uint64 `json:"id"`       // 数据ID
	GroupKey string `json:"groupKey"` // 字典分组KEY
	Label    string `json:"label"`    // 字段名称
	LabelEn  string `json:"labelEn"`  // 字段名称（英文）
	Choose   string `json:"choose"`   // 是否可被选择 0可选择 1不可选择
	Val      string `json:"val"`      // 字典值（字符型）
	Pid      uint64 `json:"pid"`      // 父级字典ID 默认 1（根数据）
	Sort     uint16 `json:"sort"`     // 字典排序
	Remark   string `json:"remark"`   // 字典描述
	Common
}

// DictTable 字典泛型造器
var DictTable actuator.Table[Dict]

// DataBase 实现指定数据库
func (t Dict) DataBase() *gorm.DB {
	return platDb
}

// TableName 实现自定义表名
func (t Dict) TableName() string {
	return "dict"
}
