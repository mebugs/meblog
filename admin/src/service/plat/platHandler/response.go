package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/plat/platService"
)

// NextResponseVal	godoc
// @id				NextResponseVal 响应码NextVal建议
// @Summary			响应码NextVal建议
// @Description		在业务分组下读取下一个响应码的建议值
// @Router			/plat/response/nextVal [post]
// @Tags			响应配置
// @Accept			json
// @Produce			json
// @Security		Token
// @Param			req	body		platModel.ResponseNextValReq	true	"请求"
// @Success			200	{object}	baseModel.ResBody{data=int64}			"响应成功"
func NextResponseVal(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &platModel.ResponseNextValReq{})
	if err == nil {
		req := reqObj.(*platModel.ResponseNextValReq)
		// 执行创建
		service.JsonRes(c, platService.NextResponseVal(traceID, req))
	}
}

// AddResponse 	godoc
// @id			AddResponse 响应码新建
// @Summary		响应码新建
// @Description	新建响应码配置
// @Router		/plat/response/add [post]
// @Tags		响应配置
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		platModel.ResponseAddReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}		"响应成功"
func AddResponse(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &platModel.ResponseAddReq{})
	if err == nil {
		req := reqObj.(*platModel.ResponseAddReq)
		// 执行创建
		service.JsonRes(c, platService.AddResponse(traceID, req))
	}
}

// PageResponse	godoc
// @id			PageResponse 响应码分页
// @Summary		响应码分页
// @Description	查询响应码分页数据
// @Router		/plat/response/page [post]
// @Tags		响应配置
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		platModel.ResponsePageReq	true											"请求"
// @Success		200	{object}	baseModel.ResBody{data=baseModel.PageRes{list=[]platModel.ResponsePageRes}}	"响应成功"
func PageResponse(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &platModel.ResponsePageReq{})
	if err == nil {
		req := reqObj.(*platModel.ResponsePageReq)
		// 执行查询
		service.JsonRes(c, platService.PageResponse(traceID, req))
	}
}

// GetResponse	godoc
// @id			GetResponse 响应码详情
// @Summary		响应码详情
// @Description	查询响应码详情
// @Router		/plat/response/get [post]
// @Tags		响应配置
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.IdReq		true							"请求"
// @Success		200	{object}	baseModel.ResBody{data=platModel.ResponseGetRes}	"响应成功"
func GetResponse(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, platService.GetResponse(traceID, req))
	}
}

// EditResponse 	godoc
// @id				EditResponse 响应码编辑
// @Summary			响应码编辑
// @Description		在响应码分组下编辑响应码
// @Router			/plat/response/edit [post]
// @Tags			响应配置
// @Accept			json
// @Produce			json
// @Security		Token
// @Param			req	body		platModel.ResponseEditReq	true	"请求"
// @Success			200	{object}	baseModel.ResBody{data=bool}		"响应成功"
func EditResponse(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &platModel.ResponseEditReq{})
	if err == nil {
		req := reqObj.(*platModel.ResponseEditReq)
		// 执行创建
		service.JsonRes(c, platService.EditResponse(traceID, req))
	}
}

// DelResponse	godoc
// @id			DelResponse 响应码封存
// @Summary		响应码封存
// @Description	响应码封存处理
// @Router		/plat/response/del [post]
// @Tags		响应配置
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		Lang	header		string				false		"语言，不传默认为zh-CN"
// @Param		req		body		[]baseModel.IdReq	true		"请求"
// @Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
func DelResponse(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, platService.DelResponse(traceID, req))
	}
}
