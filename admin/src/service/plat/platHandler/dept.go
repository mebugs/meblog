package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/plat/platService"
)

// AddDept 	godoc
// @id			AddDept 集团部门新建
// @Summary		集团部门新建
// @Description	新建集团部门
// @Router		/plat/dept/add [post]
// @Tags		集团部门
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		platModel.DeptAddReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func AddDept(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &platModel.DeptAddReq{})
	if err == nil {
		req := reqObj.(*platModel.DeptAddReq)
		// 执行创建
		service.JsonRes(c, platService.AddDept(traceID, req))
	}
}

// TreeDept	godoc
// @id			TreeDept 集团部门树
// @Summary		集团部门树
// @Description	查询集团部门树数据
// @Router		/plat/dept/tree [post]
// @Tags		访问部门
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.Req	true						"请求"
// @Success		200	{object}	baseModel.ResBody{data=[]baseModel.Tree}	"响应成功"
func TreeDept(c *gin.Context) {
	// traceID 日志追踪
	traceID := c.GetString(constant.ContextTraceID)
	// 执行查询
	service.JsonRes(c, platService.TreeDept(traceID))
}

// GetDept	godoc
// @id			GetDept 集团部门详情
// @Summary		集团部门详情
// @Description	查询集团部门详情
// @Router		/plat/dept/get [post]
// @Tags		集团部门
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.IdReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=platModel.DeptGetRes}	"响应成功"
func GetDept(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, platService.GetDept(traceID, req))
	}
}

// EditDept 	godoc
// @id			EditDept 集团部门编辑
// @Summary		集团部门编辑
// @Description	基于数据ID编辑集团部门
// @Router		/plat/dept/edit [post]
// @Tags		集团部门
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		platModel.DeptEditReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func EditDept(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &platModel.DeptEditReq{})
	if err == nil {
		req := reqObj.(*platModel.DeptEditReq)
		// 执行编辑
		service.JsonRes(c, platService.EditDept(traceID, req))
	}
}

// DelDept	godoc
// @id			DelDept 集团部门移除
// @Summary		集团部门移除
// @Description	集团部门移除处理
// @Router		/plat/dept/del [post]
// @Tags		集团部门
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req		body		baseModel.IdReq	true			"请求"
// @Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
func DelDept(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行移除
		service.JsonRes(c, platService.DelDept(traceID, req))
	}
}

// BroDept	godoc
// @id			BroDept 同级部门
// @Summary		同级部门
// @Description	同级部门列表
// @Router		/plat/dept/bro [post]
// @Tags		集团部门
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		Lang	header		string			false						"语言，不传默认为zh-CN"
// @Param		req		body		baseModel.IdReq	true						"请求"
// @Success		200		{object}	baseModel.ResBody{data=[]baseModel.SortRes}	"响应成功"
func BroDept(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, platService.BroDept(traceID, req))
	}
}

// SortDept	godoc
// @id				SortDept部门排序
// @Summary      	部门排序
// @Description  	同级部门排序功能
// @Router       	/plat/dept/sort [post]
// @Tags         	集团部门
// @Accept      	json
// @Produce      	json
// @Security	 	Token
// @Param        	req body		[]baseModel.SortReq	true	"请求"
// @Success      	200 {object}	baseModel.ResBody{data=bool}		"响应成功"
func SortDept(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &[]*baseModel.SortReq{})
	if err == nil {
		req := reqObj.(*[]*baseModel.SortReq)
		// 执行查询
		service.JsonRes(c, platService.SortDept(traceID, req))
	}
}

// ToDept	godoc
// @id				ToDept 部门迁移
// @Summary      	部门迁移
// @Description  	部门以及子部门迁移到新部门
// @Router       	/plat/dept/to [post]
// @Tags         	集团部门
// @Accept      	json
// @Produce      	json
// @Security	 	Token
// @Param        	req body		platModel.DeptToReq		true	"请求"
// @Success      	200 {object}	baseModel.ResBody{data=bool}	"响应成功"
func ToDept(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &platModel.DeptToReq{})
	if err == nil {
		req := reqObj.(*platModel.DeptToReq)
		// 执行查询
		service.JsonRes(c, platService.ToDept(traceID, req))
	}
}
