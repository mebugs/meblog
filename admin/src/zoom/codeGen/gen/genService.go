package gen

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const (
	serviceCodeTemp = `
package ${dbPack}Service

import (
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/${dbPack}Model"
	"siteol.com/smart/src/common/mysql/${dbPack}DB"
)

// Add${tableStruct} 创建${tableComment}
func Add${tableStruct}(traceID string, req *${dbPack}Model.${tableStruct}AddReq) *baseModel.ResBody {
	// 创建对象初始化
	dbReq := req.ToDbReq()
	err := ${dbPack}DB.${tableStruct}Table.InsertOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "Add${tableStruct} Fail . Err Is : %v", err)
		// 解析数据库错误
		return check${tableStruct}DBErr(err)
	}
	return baseModel.Success(constant.${tableStruct}AddSS, true)
}

// Page${tableStruct} 查询${tableComment}分页
func Page${tableStruct}(traceID string, req *${dbPack}Model.${tableStruct}PageReq) *baseModel.ResBody {
	// 查询分页
	total, list, err := ${dbPack}DB.${tableStruct}Table.Page(${tableRouter}PageQuery(req))
	if err != nil {
		log.ErrorTF(traceID, "Page${tableStruct} Fail . Err Is : %v", err)
		return baseModel.Fail(constant.${tableStruct}GetNG)
	}
	return baseModel.SuccessUnPop(baseModel.SetPageRes(${dbPack}Model.To${tableStruct}PageRes(list), total))
}

// Get${tableStruct} ${tableComment}详情
func Get${tableStruct}(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	res, err := ${dbPack}DB.${tableStruct}Table.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "Get${tableStruct} Fail . Err Is : %v", err)
		return baseModel.Fail(constant.${tableStruct}GetNG)
	}
	return baseModel.SuccessUnPop(${dbPack}Model.To${tableStruct}GetRes(&res))
}

// Edit${tableStruct} 编辑${tableComment}
func Edit${tableStruct}(traceID string, req *${dbPack}Model.${tableStruct}EditReq) *baseModel.ResBody {
	dbReq, err := ${dbPack}DB.${tableStruct}Table.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "Get${tableStruct} Fail . Err Is : %v", err)
		return baseModel.Fail(constant.${tableStruct}GetNG)
	}
	// 对象更新
	req.ToDbReq(&dbReq)
	err = ${dbPack}DB.${tableStruct}Table.UpdateOne(dbReq)
	if err != nil {
		log.ErrorTF(traceID, "Edit${tableStruct} %d Fail . Err Is : %v", dbReq.Id, err)
		// 解析数据库错误
		return check${tableStruct}DBErr(err)
	}
	return baseModel.Success(constant.${tableStruct}EditSS, true)
}

// Del${tableStruct} ${tableComment}移除
func Del${tableStruct}(traceID string, req *baseModel.IdReq) *baseModel.ResBody {
	dbReq, err := ${dbPack}DB.${tableStruct}Table.GetOneById(req.Id)
	if err != nil {
		log.ErrorTF(traceID, "Get${tableStruct} Fail . Err Is : %v", err)
		return baseModel.Fail(constant.${tableStruct}GetNG)
	}
//	// ${tableComment}禁止刪除
//	if dbReq.Mark == constant.StatusLock {
//		log.ErrorTF(traceID, "Del${tableStruct} %d Fail . Can not Edit", dbReq.Id)
//		return baseModel.Fail(constant.${tableStruct}MarkNG)
//	}
	// 物理删除
	err = ${dbPack}DB.${tableStruct}Table.DeleteOne(dbReq.Id)
	if err != nil {
		log.ErrorTF(traceID, "Del${tableStruct} %d Fail . Err Is : %v", dbReq.Id, err)
		// 硬删除直接报错
		return baseModel.Fail(constant.${tableStruct}DelNG)
	}
	return baseModel.Success(constant.${tableStruct}DelSS, true)
}

`
)

// MakeServiceCode 生成业务层类
func MakeServiceCode(tc *TableConfig, t *testing.T) error {
	// 生成代码文件
	code := strings.ReplaceAll(serviceCodeTemp, "${tableStruct}", tc.ObjName)
	code = strings.ReplaceAll(code, "${tableComment}", tc.Remark)
	code = strings.ReplaceAll(code, "${dbPack}", tc.PackName)
	code = strings.ReplaceAll(code, "${tableRouter}", tc.Router)
	// 没有目录建目录 src/service/plat/platService
	dir := fmt.Sprintf("../../service/%s/%sService", tc.PackName, tc.PackName)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		errM := os.MkdirAll(dir, 0777)
		if errM != nil {
			t.Logf("%s MakeServiceCode Mkdir Err is %v", tc.TbName, errM)
			return errM
		}
	}
	// 创建文件
	file := fmt.Sprintf("%s/%s.go", dir, tc.TbName)
	err := os.WriteFile(file, []byte(code), 0777)
	if err != nil {
		t.Logf("%s MakeServiceCode WriteFile Err is %v", tc.TbName, err)
		return err
	}
	// 执行go fmt
	cmd := exec.Command("go", "fmt", file)
	_, err = cmd.CombinedOutput()
	if err != nil {
		t.Logf("%s MakeServiceCode FMT Err is %v", tc.TbName, err)
		return err
	}
	return nil
}
