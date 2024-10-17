package gen

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const (
	serviceWorkCodeTemp = `
package ${dbPack}Service

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/${dbPack}Model"
	"siteol.com/smart/src/common/mysql/actuator"
	"strings"
)

// 业务层数据处理函数
// 抽取到独立文件中仅便于Server层阅读（没有特别意义）

// 解析数据库错误
func check${tableStruct}DBErr(err error) *baseModel.ResBody {
	errStr := err.Error()
	if strings.Contains(errStr, constant.DBDuplicateErr) {
		if strings.Contains(errStr, "xxx_uni") {
			// 唯一索引错误
			return baseModel.Fail(constant.${tableStruct}UniXxxNG)
		}
	}
	// 默认业务异常
	return baseModel.ResFail
}

// 分页查询对象封装
func ${tableRouter}PageQuery(req *platModel.${tableStruct}PageReq) (query *actuator.Query) {
	// 初始化Page
	req.PageReq.PageInit()
	// 组装Query
	query = actuator.InitQuery()
	// 模拟代码，更多函数参考Query构造器
	if req.Id != 0 {
		query.Like("id", req.Id)
	}
	if req.Id != 0 {
		query.Eq("id", req.Id)
	}
	// 模拟代码，更多函数参考Query构造器
	query.Eq("status", constant.StatusOpen)
	query.Desc("id")
	query.LimitByPage(req.Current, req.PageSize)
	return
}

`
)

// MakeServiceWorkCode 生成业务处理层类
func MakeServiceWorkCode(tc *TableConfig, t *testing.T) error {
	// 生成代码文件
	code := strings.ReplaceAll(serviceWorkCodeTemp, "${tableStruct}", tc.ObjName)
	code = strings.ReplaceAll(code, "${dbPack}", tc.PackName)
	code = strings.ReplaceAll(code, "${tableRouter}", tc.Router)
	// 没有目录建目录 src/service/plat/platService
	dir := fmt.Sprintf("../../service/%s/%sService", tc.PackName, tc.PackName)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		errM := os.MkdirAll(dir, 0777)
		if errM != nil {
			t.Logf("%s MakeServiceWorkCode Mkdir Err is %v", tc.TbName, errM)
			return errM
		}
	}
	// 创建文件
	file := fmt.Sprintf("%s/%s_work.go", dir, tc.TbName)
	err := os.WriteFile(file, []byte(code), 0777)
	if err != nil {
		t.Logf("%s MakeServiceWorkCode WriteFile Err is %v", tc.TbName, err)
		return err
	}
	// 执行go fmt
	cmd := exec.Command("go", "fmt", file)
	_, err = cmd.CombinedOutput()
	if err != nil {
		t.Logf("%s MakeServiceWorkCode FMT Err is %v", tc.TbName, err)
		return err
	}
	return nil
}
