package gen

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const (
	dbCodeTemp = `
package ${dbPack}DB

import (
	"gorm.io/gorm"
	"siteol.com/smart/src/common/mysql/actuator"
${timeImport})

// ${tableStruct} ${tableComment}
type ${tableStruct} struct {
${tableColumns}}

// ${tableStruct}Table ${tableComment}泛型造器
var ${tableStruct}Table actuator.Table[${tableStruct}]

// DataBase 实现指定数据库
func (t ${tableStruct}) DataBase() *gorm.DB {
	return ${dbPack}Db
}

// TableName 实现自定义表名
func (t ${tableStruct}) TableName() string {
	return "${table}"
}

`
)

// MakeDbCode 生成库类
func MakeDbCode(tc *TableConfig, t *testing.T) error {
	// 生成代码文件
	code := strings.ReplaceAll(dbCodeTemp, "${tableStruct}", tc.ObjName)
	code = strings.ReplaceAll(code, "${tableComment}", tc.Remark)
	code = strings.ReplaceAll(code, "${table}", tc.TbName)
	code = strings.ReplaceAll(code, "${dbPack}", tc.PackName)
	// 导入time
	timeImport := ""
	if tc.TimeImport {
		timeImport = "\"time\"\n"
	}
	code = strings.ReplaceAll(code, "${timeImport}", timeImport)
	// 组装结构体
	var sb strings.Builder
	for _, item := range tc.Columns {
		// 二维数组分别是 0 字段名 1 类型 2 JSON名 3 注释 4 非空 5 长度 6 源类型
		jsonStr := fmt.Sprintf("`json:\"%s\"`", item[2])
		sb.WriteString(fmt.Sprintf("%s %s %s // %s\n", item[0], item[1], jsonStr, item[3]))
	}
	// 表字段对象
	tableColumns := sb.String()
	code = strings.ReplaceAll(code, "${tableColumns}", tableColumns)
	// 没有目录建目录
	dir := fmt.Sprintf("../../common/mysql/%sDB", tc.PackName)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		errM := os.Mkdir(dir, 0777)
		if errM != nil {
			t.Logf("%s MakeDbCode Mkdir Err is %v", tc.TbName, errM)
			return errM
		}
	}
	// 创建文件
	file := fmt.Sprintf("%s/%s.go", dir, tc.TbName)
	err := os.WriteFile(file, []byte(code), 0777)
	if err != nil {
		t.Logf("%s MakeDbCode WriteFile Err is %v", tc.TbName, err)
		return err
	}
	// 执行go fmt
	cmd := exec.Command("go", "fmt", file)
	_, err = cmd.CombinedOutput()
	if err != nil {
		t.Logf("%s MakeDbCode FMT Err is %v", tc.TbName, err)
		return err
	}
	return nil
}
