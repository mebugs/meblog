package gen

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const (
	handlerCodeTemp = `
package ${dbPack}Handler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/${dbPack}Model"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/${dbPack}/${dbPack}Service"
)

// Add${tableStruct} 	godoc
// @id			Add${tableStruct} ${tableComment}新建
// @Summary		${tableComment}新建
// @Description	新建${tableComment}
// @Router		/${dbPack}/${tableRouter}/add [post]
// @Tags		${tableComment}
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		${dbPack}Model.${tableStruct}AddReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func Add${tableStruct}(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &${dbPack}Model.${tableStruct}AddReq{})
	if err == nil {
		req := reqObj.(*${dbPack}Model.${tableStruct}AddReq)
		// 执行创建
		service.JsonRes(c, ${dbPack}Service.Add${tableStruct}(traceID, req))
	}
}

// Page${tableStruct}	godoc
// @id			Page${tableStruct} ${tableComment}分页
// @Summary		${tableComment}分页
// @Description	分页查询${tableComment}
// @Router		/${dbPack}/${tableRouter}/page [post]
// @Tags		${tableComment}
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		${dbPack}Model.${tableStruct}PageReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=baseModel.PageRes{list=[]${dbPack}Model.${tableStruct}PageRes}}	"响应成功"
func Page${tableStruct}(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &${dbPack}Model.${tableStruct}PageReq{})
	if err == nil {
		req := reqObj.(*${dbPack}Model.${tableStruct}PageReq)
		// 执行查询
		service.JsonRes(c, ${dbPack}Service.Page${tableStruct}(traceID, req))
	}
}

// Get${tableStruct}	godoc
// @id			Get${tableStruct} ${tableComment}详情
// @Summary		${tableComment}详情
// @Description	查询${tableComment}详情
// @Router		/${dbPack}/${tableRouter}/get [post]
// @Tags		${tableComment}
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.IdReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=${dbPack}Model.${tableStruct}GetRes}	"响应成功"
func Get${tableStruct}(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, ${dbPack}Service.Get${tableStruct}(traceID, req))
	}
}

// Edit${tableStruct} 	godoc
// @id			Edit${tableStruct} ${tableComment}编辑
// @Summary		${tableComment}编辑
// @Description	基于数据ID编辑${tableComment}
// @Router		/${dbPack}/${tableRouter}/edit [post]
// @Tags		${tableComment}
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		${dbPack}Model.${tableStruct}EditReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func Edit${tableStruct}(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &${dbPack}Model.${tableStruct}EditReq{})
	if err == nil {
		req := reqObj.(*${dbPack}Model.${tableStruct}EditReq)
		// 执行编辑
		service.JsonRes(c, ${dbPack}Service.Edit${tableStruct}(traceID, req))
	}
}

// Del${tableStruct}	godoc
// @id			Del${tableStruct} ${tableComment}移除
// @Summary		${tableComment}移除
// @Description	${tableComment}移除处理
// @Router		/${dbPack}/${tableRouter}/del [post]
// @Tags		${tableComment}
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req		body		baseModel.IdReq	true			"请求"
// @Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
func Del${tableStruct}(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行移除
		service.JsonRes(c, ${dbPack}Service.Del${tableStruct}(traceID, req))
	}
}

`
)

// MakeHandlerCode 生成控制层类
func MakeHandlerCode(tc *TableConfig, t *testing.T) error {
	// 生成代码文件
	code := strings.ReplaceAll(handlerCodeTemp, "${tableStruct}", tc.ObjName)
	code = strings.ReplaceAll(code, "${tableComment}", tc.Remark)
	code = strings.ReplaceAll(code, "${dbPack}", tc.PackName)
	code = strings.ReplaceAll(code, "${tableRouter}", tc.Router)
	// 没有目录建目录 src/service/plat/platHandler
	dir := fmt.Sprintf("../../service/%s/%sHandler", tc.PackName, tc.PackName)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		errM := os.MkdirAll(dir, 0777)
		if errM != nil {
			t.Logf("%s MakeHandlerCode Mkdir Err is %v", tc.TbName, errM)
			return errM
		}
	}
	// 创建文件
	file := fmt.Sprintf("%s/%s.go", dir, tc.TbName)
	err := os.WriteFile(file, []byte(code), 0777)
	if err != nil {
		t.Logf("%s MakeHandlerCode WriteFile Err is %v", tc.TbName, err)
		return err
	}
	// 执行go fmt
	cmd := exec.Command("go", "fmt", file)
	_, err = cmd.CombinedOutput()
	if err != nil {
		t.Logf("%s MakeHandlerCode FMT Err is %v", tc.TbName, err)
		return err
	}
	return nil
}
