package respCode

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

/**
 * 基于Excel文件生成响应码国际化
 * 1. 生成用于constant包下的response_code.go（生成后利用 go fmt完成格式化）
 * 2. 生成刷库SQL
 * 3. 执行目录为当前目录
 */

// TestMakeResp
func TestMakeResp(t *testing.T) {
	// 打开表格文件
	excel, err := excelize.OpenFile("response_code.xlsx")
	if err != nil {
		t.Logf("Open response_code.xlsx  Fail . Err Is %v", err)
		return
	}
	// 读取指定工作表所有数据
	rows, err := excel.GetRows("response_code")
	if err != nil {
		t.Logf("Read response_code Data Fail . Err Is %v", err)
		return
	}
	t.Logf("Read response_code Success")
	// SQL Array
	sqlBuilder := &strings.Builder{}
	// Go File
	goBuilder := &strings.Builder{}
	// 添加开头内容
	sqlBuilder.WriteString(`SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
-- Reset Table
TRUNCATE TABLE response_code;
-- Record Start
`)
	goBuilder.WriteString(`package constant

/**
 *
 * 响应码常量
 * 具体文言维护在数据库，但编码需要在此处维护，提高代码可读性
 *
 * @author 米虫丨www.mebugs.com
 * @since 2023-07-21
 */

const (
	CodeFmt = "%s%s%02d" // 成功码Format
`)
	// 添加实际内容
	for i, data := range rows {
		if i == 0 {
			continue
		}
		// 写入builder
		// 代码划区
		if data[3] != rows[i-1][3] {
			goBuilder.WriteString("\n")
		}
		// ID 响应码拼接
		id, _ := strconv.Atoi(data[0])
		codeSeq, _ := strconv.Atoi(data[4])
		code := fmt.Sprintf("%s%s%02d", data[2], data[3], codeSeq)
		// SQL处理
		sql := fmt.Sprintf("INSERT INTO response_code VALUES ('%d', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '0', null, null);\n", id, code, data[3], data[2], data[7], data[8], data[6], data[5])
		sqlBuilder.WriteString(sql)
		// 代码处理
		goBuilder.WriteString(fmt.Sprintf("	%s  = \"%s\" // %s\n", data[1], code, data[6]))
	}
	// 添加结尾内容
	sqlBuilder.WriteString(`-- Record End
SET FOREIGN_KEY_CHECKS = 1;
`)
	goBuilder.WriteString(`
)
`)
	err = os.WriteFile("response_code.sql", []byte(sqlBuilder.String()), 0777)
	if err != nil {
		t.Logf("Write response_code.sql Fail . Err Is %v ", err)
		return
	}
	t.Logf("Write response_code.sql Success")
	// 代码覆盖
	err = os.WriteFile("../../common/constant/response_code.go", []byte(goBuilder.String()), 0777)
	if err != nil {
		t.Logf("Write response_code.go Fail . Err Is %v ", err)
		return
	}
	t.Logf("Write response_code.go Success")
	// 执行go fmt
	cmd := exec.Command("go", "fmt", "../../common/constant/response_code.go")
	_, err = cmd.CombinedOutput()
	if err != nil {
		t.Logf("Go FMT Fail %v", err)
		return
	}
	t.Logf("Go FMT response_code.go Success")
	t.Logf("Done!!!")
}
