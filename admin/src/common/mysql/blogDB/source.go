package blogDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
	"time"
)

// Source 资源配置表
type Source struct {
	Id       uint64     `json:"id"`       // 数据ID
	Name     string     `json:"name"`     // 资源名
	BackEnd  string     `json:"backEnd"`  // 图片后缀，jpg png gif
	FileType uint8      `json:"fileType"` // 图片类型，枚举：0_主图 1_图标 2_正文
	FilePath string     `json:"filePath"` // 文件子目录，年份目录
	Status   string     `json:"status"`   // 状态，枚举：0_正常 1_锁定 2_封存
	CreateAt *time.Time `json:"createAt"` // 创建时间
	UpdateAt *time.Time `json:"updateAt"` // 更新时间
}

// SourceTable 资源配置表泛型造器
var SourceTable actuator.Table[Source]

// DataBase 实现指定数据库
func (t Source) DataBase() *gorm.DB {
	return blogDb
}

// TableName 实现自定义表名
func (t Source) TableName() string {
	return "source"
}
