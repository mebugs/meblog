package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/plat/platService"
)

// AddRole	godoc
// @id				AddRole 角色新建
// @Summary			角色新建
// @Description		新建角色配置
// @Router			/plat/role/add [post]
// @Tags			角色配置
// @Accept			json
// @Produce			json
// @Security		Token
// @Param			req	body		platModel.RoleAddReq	true	"请求"
// @Success			200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func AddRole(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &platModel.RoleAddReq{})
	if err == nil {
		req := reqObj.(*platModel.RoleAddReq)
		// 执行创建
		service.JsonRes(c, platService.AddRole(traceID, req))
	}
}

// PageRole	godoc
// @id			PageRole 角色分页
// @Summary		角色分页
// @Description	查询角色分页数据
// @Router		/plat/role/page [post]
// @Tags		角色配置
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		platModel.RolePageReq	true											"请求"
// @Success		200	{object}	baseModel.ResBody{data=baseModel.PageRes{list=[]platModel.RolePageRes}}	"响应成功"
func PageRole(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &platModel.RolePageReq{})
	if err == nil {
		req := reqObj.(*platModel.RolePageReq)
		// 执行查询
		service.JsonRes(c, platService.PageRole(traceID, req))
	}
}

// GetRole	godoc
// @id			GetRole 角色详情
// @Summary		角色详情
// @Description	查询角色详情
// @Router		/plat/role/get [post]
// @Tags		角色配置
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.IdReq		true						"请求"
// @Success		200	{object}	baseModel.ResBody{data=platModel.RoleGetRes}	"响应成功"
func GetRole(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, platService.GetRole(traceID, req))
	}
}

// EditRole 	godoc
// @id				EditRole 角色编辑
// @Summary			角色编辑
// @Description		在角色分组下编辑角色
// @Router			/plat/role/edit [post]
// @Tags			角色配置
// @Accept			json
// @Produce			json
// @Security		Token
// @Param			req	body		platModel.RoleEditReq	true	"请求"
// @Success			200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func EditRole(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &platModel.RoleEditReq{})
	if err == nil {
		req := reqObj.(*platModel.RoleEditReq)
		// 执行创建
		service.JsonRes(c, platService.EditRole(traceID, req))
	}
}

// DelRole	godoc
// @id			DelRole 角色删除
// @Summary		角色删除
// @Description	角色删除处理
// @Router		/plat/role/del [post]
// @Tags		角色配置
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		Lang	header		string				false		"语言，不传默认为zh-CN"
// @Param		req		body		baseModel.IdReq		true		"请求"
// @Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
func DelRole(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, platService.DelRole(traceID, req))
	}
}

// ListRole	godoc
// @id			ListRole 角色下拉列表
// @Summary		角色列表
// @Description	角色下拉列表
// @Router		/plat/role/list [post]
// @Tags		角色配置
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		Lang	header		string				false							"语言，不传默认为zh-CN"
// @Success		200		{object}	baseModel.ResBody{data=[]baseModel.SelectNumRes}	"响应成功"
func ListRole(c *gin.Context) {
	traceID := c.GetString(constant.ContextTraceID)
	service.JsonRes(c, platService.ListRole(traceID))
}
