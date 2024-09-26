package platDB

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

const (
	temp = `
package platDB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
${imports}
)

// ${tableStruct} ${tableComment}
type ${tableStruct} struct {
${tableColumns}}

// ${tableStruct}Table ${tableComment}泛型造器
var ${tableStruct}Table actuator.Table[${tableStruct}]

// DataBase 实现指定数据库
func (t ${tableStruct}) DataBase() *gorm.DB {
	return platDB
}

// TableName 实现自定义表名
func (t ${tableStruct}) TableName() string {
	return "${table}"
}

`
)

var typeMap = map[string]string{
	"bigint":   "uint64",
	"int":      "uint16",
	"smallint": "uint8",
	"varchar":  "string",
	"text":     "string",
	"datetime": "*time.Time",
}
var commonColumn = map[string]bool{"mark": true, "status": true, "create_at": true, "update_at": true}

// 计划生成的表名 - 平台表
var tableArray = []string{
	//"account",
	//"account_role",
	//"dept",
	//"dict",
	//"dict_group",
	//"login_record",
	//"permission",
	//"permission_router",
	//"response_code",
	//"role",
	//"role_permission",
	//"router",
	//"sys_config",
	"router_log",
}

// 包位置
var packPath = "platDB"

// 库名
var dbName = "smart"

func TestDBPlatMaker(t *testing.T) {
	InitPlatFromDb("Test")
	path := fmt.Sprintf("src/common/mysql/%s/", packPath)
	// 切换到管理库
	platDb.Exec("use information_schema;")
	for _, table := range tableArray {
		// 表结构体 表文件 表注释
		var tableStruct, tableFile, tableComment string
		tableStruct = toFixStr(table, true)
		tableFile = table + ".go"
		tx := platDb.Raw(fmt.Sprintf("select `table_comment` from `tables` where `table_schema` = '%s' AND `table_name`='%s'", dbName, table)).Scan(&tableComment)
		if tableComment == "" {
			t.Logf("Make %s File Fail .Can not Get %s , Err Is : %v", tableFile, table, tx.Error)
			continue
		}
		// 是否需要追加导入 是否包含公共字段
		var importHave, commonHave bool
		// 字段和类型的最大占位（自动Format）
		var aL, bL int
		// 用于引入代码文件的字段列表 字段名 类型 JSON名 注释
		columnList := make([][4]string, 0)
		rows, err := platDb.Raw(fmt.Sprintf("select `column_name`,`data_type`,`column_comment` from `columns` where `table_schema` = '%s' AND `table_name`='%s' ORDER BY ORDINAL_POSITION", dbName, table)).Rows()
		if err != nil {
			t.Logf("Make %s File Fail .Can not Get %s Columns , Err Is : %v", tableFile, table, err)
			continue
		}
		// 组装字段
		for rows.Next() {
			var a, b, c string
			_ = rows.Scan(&a, &b, &c)
			// 公共字段
			if commonColumn[a] {
				commonHave = true
				if aL < 6 {
					aL = 6
				}
			} else {
				if len(a) > aL {
					aL = len(a)
				}
				if b == "datetime" {
					importHave = true
				}
				bs := typeMap[b]
				if len(bs) > bL {
					bL = len(bs)
				}
				columnList = append(columnList, [4]string{toFixStr(a, true), bs, toFixStr(a, false), c})
			}
		}
		// 组装结构体
		var sb strings.Builder
		for _, item := range columnList {
			tab1 := strings.Repeat(" ", aL-len(item[0]))
			tab2 := strings.Repeat(" ", bL-len(item[1]))
			jsonStr := fmt.Sprintf("`json:\"%s\"`", item[2])
			sb.WriteString(fmt.Sprintf("\t%s %s %s %s %s%s // %s\n", item[0], tab1, item[1], tab2, jsonStr, tab1, item[3]))
		}
		if commonHave {
			sb.WriteString("\tCommon\n")
		}
		// 表字段对象
		tableColumns := sb.String()
		imports := ""
		if importHave {
			imports = "\t\"time\""
		}
		// 生成代码文件
		code := strings.ReplaceAll(temp, "${imports}", imports)
		code = strings.ReplaceAll(code, "${tableStruct}", tableStruct)
		code = strings.ReplaceAll(code, "${tableComment}", tableComment)
		code = strings.ReplaceAll(code, "${tableColumns}", tableColumns)
		code = strings.ReplaceAll(code, "${table}", table)
		err = os.WriteFile(path+tableFile, []byte(code), 0777)
		if err != nil {
			t.Logf("Make %s File Fail . Err Is : %v", tableFile, err)
		} else {
			t.Logf("Make %s File Success . ", tableFile)
		}
	}
}

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
