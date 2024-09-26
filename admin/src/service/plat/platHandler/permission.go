package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/plat/platService"
)

// AddPermission	godoc
// @id				AddPermission 权限新建
// @Summary			权限新建
// @Description		新建权限配置
// @Router			/plat/permission/add [post]
// @Tags			访问权限
// @Accept			json
// @Produce			json
// @Security		Token
// @Param			req	body		platModel.PermissionAddReq	true	"请求"
// @Success			200	{object}	baseModel.ResBody{data=bool}		"响应成功"
func AddPermission(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &platModel.PermissionAddReq{})
	if err == nil {
		req := reqObj.(*platModel.PermissionAddReq)
		// 执行创建
		service.JsonRes(c, platService.AddPermission(traceID, req))
	}
}

// TreePermission	godoc
// @id			PagePermission 权限树
// @Summary		权限树
// @Description	查询权限树数据
// @Router		/plat/permission/tree [post]
// @Tags		访问权限
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.Req	true						"请求"
// @Success		200	{object}	baseModel.ResBody{data=[]baseModel.Tree}	"响应成功"
func TreePermission(c *gin.Context) {
	// traceID 日志追踪
	traceID := c.GetString(constant.ContextTraceID)
	// 执行查询
	service.JsonRes(c, platService.TreePermission(traceID))
}

// GetPermission	godoc
// @id			GetPermission 权限详情
// @Summary		权限详情
// @Description	查询权限详情
// @Router		/plat/permission/get [post]
// @Tags		访问权限
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.IdReq		true							"请求"
// @Success		200	{object}	baseModel.ResBody{data=platModel.PermissionGetRes}	"响应成功"
func GetPermission(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, platService.GetPermission(traceID, req))
	}
}

// EditPermission 	godoc
// @id				EditPermission 权限编辑
// @Summary			权限编辑
// @Description		在权限分组下编辑权限
// @Router			/plat/permission/edit [post]
// @Tags			访问权限
// @Accept			json
// @Produce			json
// @Security		Token
// @Param			req	body		platModel.PermissionEditReq	true	"请求"
// @Success			200	{object}	baseModel.ResBody{data=bool}		"响应成功"
func EditPermission(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &platModel.PermissionEditReq{})
	if err == nil {
		req := reqObj.(*platModel.PermissionEditReq)
		// 执行创建
		service.JsonRes(c, platService.EditPermission(traceID, req))
	}
}

// DelPermission	godoc
// @id			DelPermission 权限封存
// @Summary		权限封存
// @Description	权限封存处理
// @Router		/plat/permission/del [post]
// @Tags		访问权限
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		Lang	header		string				false		"语言，不传默认为zh-CN"
// @Param		req		body		baseModel.IdReq		true		"请求"
// @Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
func DelPermission(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, platService.DelPermission(traceID, req))
	}
}

// BroPermission	godoc
// @id			BroPermission 同级权限
// @Summary		同级权限
// @Description	同级权限列表
// @Router		/plat/permission/bro [post]
// @Tags		访问权限
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		Lang	header		string			false						"语言，不传默认为zh-CN"
// @Param		req		body		baseModel.IdReq	true						"请求"
// @Success		200		{object}	baseModel.ResBody{data=[]baseModel.SortRes}	"响应成功"
func BroPermission(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, platService.BroPermission(traceID, req))
	}
}

// SortPermission	godoc
// @id				SortPermission权限排序
// @Summary      	权限排序
// @Description  	同级权限排序功能
// @Router       	/plat/permission/sort [post]
// @Tags         	访问权限
// @Accept      	json
// @Produce      	json
// @Security	 	Token
// @Param        	req body		[]baseModel.SortReq	true	"请求"
// @Success      	200 {object}	baseModel.ResBody{data=bool}		"响应成功"
func SortPermission(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &[]*baseModel.SortReq{})
	if err == nil {
		req := reqObj.(*[]*baseModel.SortReq)
		// 执行查询
		service.JsonRes(c, platService.SortPermission(traceID, req))
	}
}
