package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/plat/platService"
)

// AddRouter 	godoc
// @id			AddRouter 路由接口新建
// @Summary		路由接口新建
// @Description	新建路由接口配置
// @Router		/plat/router/add [post]
// @Tags		路由接口
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		platModel.RouterAddReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func AddRouter(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &platModel.RouterAddReq{})
	if err == nil {
		req := reqObj.(*platModel.RouterAddReq)
		// 执行创建
		service.JsonRes(c, platService.AddRouter(traceID, req))
	}
}

// PageRouter	godoc
// @id			PageRouter 路由接口分页
// @Summary		路由接口分页
// @Description	查询路由接口分页数据
// @Router		/plat/router/page [post]
// @Tags		路由接口
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		platModel.RouterPageReq	true												"请求"
// @Success		200	{object}	baseModel.ResBody{data=baseModel.PageRes{list=[]platModel.RouterPageRes}}	"响应成功"
func PageRouter(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &platModel.RouterPageReq{})
	if err == nil {
		req := reqObj.(*platModel.RouterPageReq)
		// 执行查询
		service.JsonRes(c, platService.PageRouter(traceID, req))
	}
}

// GetRouter	godoc
// @id			GetRouter 路由接口详情
// @Summary		路由接口详情
// @Description	查询路由接口详情
// @Router		/plat/router/get [post]
// @Tags		路由接口
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.IdReq		true						"请求"
// @Success		200	{object}	baseModel.ResBody{data=platModel.RouterGetRes}	"响应成功"
func GetRouter(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, platService.GetRouter(traceID, req))
	}
}

// EditRouter 	godoc
// @id				EditRouter 路由接口编辑
// @Summary			路由接口编辑
// @Description		在路由接口分组下编辑路由接口
// @Router			/plat/router/edit [post]
// @Tags			路由接口
// @Accept			json
// @Produce			json
// @Security		Token
// @Param			req	body		platModel.RouterEditReq	true	"请求"
// @Success			200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func EditRouter(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &platModel.RouterEditReq{})
	if err == nil {
		req := reqObj.(*platModel.RouterEditReq)
		// 执行创建
		service.JsonRes(c, platService.EditRouter(traceID, req))
	}
}

// DelRouter	godoc
// @id			DelRouter 路由接口封存
// @Summary		路由接口封存
// @Description	路由接口封存处理
// @Router		/plat/router/del [post]
// @Tags		路由接口
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		Lang	header		string				false		"语言，不传默认为zh-CN"
// @Param		req		body		baseModel.IdReq	true		"请求"
// @Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
func DelRouter(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, platService.DelRouter(traceID, req))
	}
}
