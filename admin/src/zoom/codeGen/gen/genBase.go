package gen

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

var db *gorm.DB

var typeMap = map[string]string{
	"bigint":   "uint64",
	"int":      "uint16",
	"smallint": "uint8",
	"varchar":  "string",
	"text":     "string",
	"longtext": "string",
	"datetime": "*time.Time",
}

// InitDb 初始化数据库连接
func InitDb(url string) error {
	// 采用默认配置打开数据可（默认禁用事务）
	dbInit, err := gorm.Open(mysql.Open(url), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		return err
	}
	db = dbInit
	return nil
}

// GetTableBase 读取表注释
func GetTableBase(packName, dbName, tbName string) (tc *TableConfig, err error) {
	var tableComment string
	tx := db.Raw(fmt.Sprintf("select `table_comment` from `tables` where `table_schema` = '%s' AND `table_name`='%s'", dbName, tbName)).Scan(&tableComment)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	if tableComment == "" {
		err = errors.New("CommentIsEmpty")
		return
	}
	tc = &TableConfig{
		PackName: packName,
		DbName:   dbName,
		TbName:   tbName,
		ObjName:  toFixStr(tbName, true),
		Router:   toFixStr(tbName, false),
		Remark:   tableComment,
		Columns:  nil,
	}
	return
}

// GetTableColumns 读取表的字段
func GetTableColumns(tc *TableConfig) error {
	// 用于引入代码文件的字段列表
	columnList := make([][7]string, 0)
	baseSql := "select `column_name`,`data_type`,`column_comment`,`IS_NULLABLE`,`CHARACTER_MAXIMUM_LENGTH` from `columns` where `table_schema` = ? AND `table_name`= ? ORDER BY ORDINAL_POSITION"
	rows, err := db.Raw(baseSql, tc.DbName, tc.TbName).Rows()
	if err != nil {
		return err
	}
	// 组装字段
	for rows.Next() {
		var a, b, c, d string
		var e sql.NullInt64
		err = rows.Scan(&a, &b, &c, &d, &e)
		if err != nil {
			return err
		}
		eStr := fmt.Sprintf("%d", e.Int64)
		structType := typeMap[b]
		if structType == "*time.Time" {
			tc.TimeImport = true
		}
		// 二维数组分别是 0 字段名 1 类型 2 JSON名 3 注释 4 非空 5 长度 6 源类型
		columnList = append(columnList, [7]string{toFixStr(a, true), structType, toFixStr(a, false), c, d, eStr, b})
	}
	tc.Columns = columnList
	return nil
}

// 处理文件名 结构体名 pre首字母大写是否
func toFixStr(s string, pre bool) string {
	// 首字母大写
	ss := s
	if pre {
		ss = strings.ToUpper(s[0:1]) + s[1:]
	}
	sn := ""
	// 下换线转换 _x => X
	ci := -1
	for i, si := range ss {
		if string(si) == "_" {
			ci = i + 1
		} else {
			if i == ci {
				sn = sn + strings.ToUpper(string(si))
			} else {
				sn = sn + string(si)
			}
		}
	}
	return sn
}
