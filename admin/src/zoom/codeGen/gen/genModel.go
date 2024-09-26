package gen

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const (
	modelCodeTemp = `
package platModel

import (
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/mysql/${dbPack}DB"
${timeImport})

// ${tableStruct}DoReq ${tableComment} 通用请求，创建&编辑可复用的字段
type ${tableStruct}DoReq struct {
	${tableColumnsWithOutId}}

// ${tableStruct}AddReq ${tableComment} 创建请求，酌情从通用中摘出部分字段
type ${tableStruct}AddReq struct {
	${tableStruct}DoReq
}

// ${tableStruct}EditReq ${tableComment} 编辑请求，酌情从通用中摘出部分字段
type ${tableStruct}EditReq struct {
	${tableColumnsId} 
	${tableStruct}DoReq
}

// ToDbReq ${tableComment} 创建转数据库
func (r *${tableStruct}AddReq) ToDbReq() *${dbPack}DB.${tableStruct} {
	${hasTime}return &platDB.${tableStruct}{
		Id:           0,
		${tableColumnsToAdd}}
}

// ToDbReq ${tableComment} 更新转数据库
func (r *${tableStruct}EditReq) ToDbReq(d *platDB.${tableStruct}) {
	${hasTime}${tableColumnsToMod}}

// ${tableStruct}GetRes ${tableComment} 详情响应
type ${tableStruct}GetRes struct {
	${tableColumnsIdUnBinding} 
	${tableColumnsUnBinding}}

// ${tableStruct}PageReq ${tableComment} 分页请求，根据实际业务替换分页条件字段
type ${tableStruct}PageReq struct {
	${tableColumnsIdUnBinding} 
	baseModel.PageReq
}

// ${tableStruct}PageRes ${tableComment} 分页响应，酌情从详情摘出部分字段
type ${tableStruct}PageRes struct {
	${tableStruct}GetRes
}

// To${tableStruct}GetRes ${tableComment} 数据库转为详情响应
func To${tableStruct}GetRes(r *platDB.${tableStruct}) *${tableStruct}GetRes {
	return &${tableStruct}GetRes{
		Id:	r.Id,
		${tableColumnsGetRes}}}

// To${tableStruct}PageRes ${tableComment} 数据库转分页响应
func To${tableStruct}PageRes(list []*platDB.${tableStruct}) []*${tableStruct}PageRes {
	res := make([]*${tableStruct}PageRes, len(list))
	for i, r := range list {
		res[i] = &${tableStruct}PageRes{
			${tableStruct}GetRes: *To${tableStruct}GetRes(r),
		}
	}
	return res
}

`
	modelCodeIdLine          = "Id uint64 `json:\"id\" binding:\"required\" example:\"1\"` // 数据ID"
	modelCodeIdLineUnBinding = "Id uint64 `json:\"id\" example:\"1\"` // 数据ID"
)

// MakeModelCode 生成库类
func MakeModelCode(tc *TableConfig, t *testing.T) error {
	var cs strings.Builder
	var csUnBind strings.Builder
	var cToAdd strings.Builder
	var cToMod strings.Builder
	var cToRes strings.Builder
	var hasTime string
	if tc.TimeImport {
		hasTime = "now := time.Now()\n"
	}
	for _, item := range tc.Columns {
		if item[0] == "Id" {
			continue
		}
		// 生成校验字串
		bindStr, exampleStr := makeBindAndExample(item)
		// 组成字段串
		// 二维数组分别是 0 字段名 1 类型 2 JSON名 3 注释 4 非空 5 长度 6 源类型
		cs.WriteString(fmt.Sprintf("%s %s `json:\"%s\"%s%s`// %s\n", item[0], item[1], item[2], bindStr, exampleStr, item[3]))
		csUnBind.WriteString(fmt.Sprintf("%s %s `json:\"%s\"%s`// %s\n", item[0], item[1], item[2], exampleStr, item[3]))
		if item[1] == "*time.Time" {
			cToAdd.WriteString(fmt.Sprintf("%s: &now,\n", item[0]))
			cToMod.WriteString(fmt.Sprintf("d.%s = &now\n", item[0]))
		} else {
			cToAdd.WriteString(fmt.Sprintf("%s:r.%s,\n", item[0], item[0]))
			cToMod.WriteString(fmt.Sprintf("d.%s = r.%s\n", item[0], item[0]))
		}
		cToRes.WriteString(fmt.Sprintf("%s:r.%s,\n", item[0], item[0]))
	}
	// 处理模版
	code := strings.ReplaceAll(modelCodeTemp, "${tableStruct}", tc.ObjName)
	code = strings.ReplaceAll(code, "${tableComment}", tc.Remark)
	code = strings.ReplaceAll(code, "${table}", tc.TbName)
	code = strings.ReplaceAll(code, "${dbPack}", tc.PackName)
	code = strings.ReplaceAll(code, "${hasTime}", hasTime)
	code = strings.ReplaceAll(code, "${tableColumnsIdUnBinding}", modelCodeIdLineUnBinding)
	code = strings.ReplaceAll(code, "${tableColumnsId}", modelCodeIdLine)
	code = strings.ReplaceAll(code, "${tableColumnsWithOutId}", cs.String())
	code = strings.ReplaceAll(code, "${tableColumnsToAdd}", cToAdd.String())
	code = strings.ReplaceAll(code, "${tableColumnsToMod}", cToMod.String())
	code = strings.ReplaceAll(code, "${tableColumnsUnBinding}", csUnBind.String())
	code = strings.ReplaceAll(code, "${tableColumnsGetRes}", cToRes.String())
	// 导入time
	timeImport := ""
	if tc.TimeImport {
		timeImport = "\"time\"\n"
	}
	code = strings.ReplaceAll(code, "${timeImport}", timeImport)
	// 没有目录建目录
	dir := fmt.Sprintf("../../common/model/%sModel", tc.PackName)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		errM := os.Mkdir(dir, 0777)
		if errM != nil {
			t.Logf("%s MakeModelCode Mkdir Err is %v", tc.TbName, errM)
			return errM
		}
	}
	// 创建文件
	file := fmt.Sprintf("%s/%s.go", dir, tc.TbName)
	err := os.WriteFile(file, []byte(code), 0777)
	if err != nil {
		t.Logf("%s MakeModelCode WriteFile Err is %v", tc.TbName, err)
		return err
	}
	// 执行go fmt
	cmd := exec.Command("go", "fmt", file)
	_, err = cmd.CombinedOutput()
	if err != nil {
		t.Logf("%s MakeModelCode FMT Err is %v", tc.TbName, err)
		return err
	}
	return nil
}

// 生成校验字串和样例
func makeBindAndExample(item [7]string) (string, string) {
	// 二维数组分别是 0 字段名 1 类型 2 JSON名 3 注释 4 非空 5 长度 6 源类型
	var need, length, oneOf bool
	// 样例字串
	example := ""
	switch item[1] {
	case "uint64", "uint16", "uint8":
		example = "0"
	case "string":
		example = "demo"
	}
	// 数据库必填，默认接口必传，可以服务端赋值
	if item[4] == "NO" {
		need = true
	}
	// 字符串类型可能需要配置长度校验
	if item[6] == "varchar" {
		length = true
	}
	// 枚举情况下，无需长度限制
	if strings.Contains(item[3], "枚举：") {
		oneOf = true
	}
	var h bool
	var bind strings.Builder
	if need {
		h = true
		bind.WriteString("required")
	}
	if length && !oneOf {
		if h {
			bind.WriteString(",")
		}
		bind.WriteString(fmt.Sprintf("max=%s", item[5]))
	}
	if oneOf {
		// 枚举处理 状态，枚举：0_正常 1_锁定 2_封存
		mA := make([]string, 0)
		cA := strings.Split(item[3], "枚举：")
		if len(cA) == 2 {
			// 0_正常
			cM := strings.Split(cA[1], " ")
			if len(cM) > 0 {
				for _, cMI := range cM {
					cMIA := strings.Split(cMI, "_")
					// 0 正常
					if len(cMIA) == 2 {
						mA = append(mA, cMIA[0])
					}
				}
			}
		}
		if len(mA) > 0 {
			if h {
				bind.WriteString(",")
			}
			bind.WriteString("oneof=")
			example = mA[0]
			for i, m := range mA {
				if i > 0 {
					bind.WriteString(" ")
				}
				bind.WriteString(fmt.Sprintf("'%s'", m))
			}
		}
	}
	// 校验字串
	bindStr := ""
	if bind.Len() > 0 {
		bindStr = fmt.Sprintf(" binding:\"%s\"", bind.String())
	}
	// 样例字段
	exampleStr := ""
	if example != "" {
		exampleStr = fmt.Sprintf(" example:\"%s\"", example)
	}
	return bindStr, exampleStr
}
