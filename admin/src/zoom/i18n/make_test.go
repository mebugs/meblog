package i18n

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strings"
	"testing"
)

/**
 * 基于Excel文件生成多语言国际化
 * Sheet名是国际化文件名
 * 从第二格开始表头是语言名
 * 编译成二进制文件使用 go test -c
 */

// TestMakeI18N 二进制文件执行方式
func TestMakeI18N(t *testing.T) {

	// 打开表格文件
	excel, err := excelize.OpenFile("i18n.xlsx")
	if err != nil {
		t.Logf("Open I18n.xlsx  Fail . Err Is %v", err)
		return
	}

	// 读取指定工作表所有数据
	langRows, err := excel.GetRows("lang")
	if err != nil {
		t.Logf("Read Lang Config Fail . Err Is %v", err)
		return
	}
	// 全部语言
	allLang := make([]string, len(langRows)-1)
	for i, row := range langRows {
		if i == 0 {
			continue
		}
		if len(row) > 0 {
			allLang[i-1] = row[0]
			// 创建语言目录
			if _, errS := os.Stat(row[0]); os.IsNotExist(errS) {
				errM := os.Mkdir(row[0], 0777)
				if errM != nil {
					t.Logf("Make %s Dire Fail . Err Is %v", row[0], errM)
					return
				}
			}
		}
	}
	t.Logf("Set All Lang Is %s", allLang)
	// 查看全部的sheet
	sheets := excel.GetSheetList()
	// 遍历全部表
	for _, sheet := range sheets {
		if sheet == "lang" {
			continue
		}
		// 读取语言数据
		dataRows, errG := excel.GetRows(sheet)
		if errG != nil {
			t.Logf("Read %s Fail . Err Is %v ", sheet, errG)
			return
		}
		// 创建语言对象
		dataArray := make([]*strings.Builder, len(allLang))
		for i, _ := range dataArray {
			dataArray[i] = &strings.Builder{}
			dataArray[i].WriteString("export default {\n")
		}
		// 循环处理数据
		for i, data := range dataRows {
			if i == 0 {
				continue
			}
			// 写入builder
			for c, b := range dataArray {
				b.WriteString(fmt.Sprintf("  '%s': '%s',\n", data[0], data[c+1]))
			}
		}
		for i, _ := range dataArray {
			dataArray[i].WriteString("}\n")
			// 写出文件
			fileName := allLang[i] + "/" + sheet + ".ts"
			errW := os.WriteFile(fileName, []byte(dataArray[i].String()), 0777)
			if errW != nil {
				t.Logf("Write %s Fail . Err Is %v ", fileName, errW)
				return
			}
			t.Logf("Write %s Success .", fileName)
		}
	}
	t.Logf("Done!!!")
}
