package gen

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const (
	serviceConstantCodeTemp = `
	package constant

/**
 * 请将下述常量迁移到 response_code.go
 * 相关编号通过 src/zoom/respCode/make_test.go 生成
 */

const (
	${tableStruct}AddSS    = "Sn00" // ${tableComment}创建成功
	${tableStruct}EditSS   = "Sn01" // ${tableComment}编辑成功
	${tableStruct}DelSS    = "Sn02" // ${tableComment}删除成功
	${tableStruct}GetNG    = "Fn00" // ${tableComment}查询失败
	${tableStruct}UniXxxNG = "Fn01" // ${tableComment}地址全局唯一
	${tableStruct}MarkNG   = "Fn02" // 内置${tableComment}禁止删除
	${tableStruct}DelNG    = "Fn03" // ${tableComment}删除
	
)

`
)

// MakeConstantDemoCode 生成常量Demo类
func MakeConstantDemoCode(tc *TableConfig, t *testing.T) error {
	// 生成代码文件
	code := strings.ReplaceAll(serviceConstantCodeTemp, "${tableStruct}", tc.ObjName)
	code = strings.ReplaceAll(code, "${tableComment}", tc.Remark)
	// 没有目录建目录 src/common/constant
	dir := fmt.Sprintf("../../common/constant")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		errM := os.Mkdir(dir, 0777)
		if errM != nil {
			t.Logf("%s MakeConstantDemoCode Mkdir Err is %v", tc.TbName, errM)
			return errM
		}
	}
	// 创建文件
	file := fmt.Sprintf("%s/demo_constant_%s.go", dir, tc.TbName)
	err := os.WriteFile(file, []byte(code), 0777)
	if err != nil {
		t.Logf("%s MakeConstantDemoCode WriteFile Err is %v", tc.TbName, err)
		return err
	}
	// 执行go fmt
	cmd := exec.Command("go", "fmt", file)
	_, err = cmd.CombinedOutput()
	if err != nil {
		t.Logf("%s MakeConstantDemoCode FMT Err is %v", tc.TbName, err)
		return err
	}
	return nil
}
