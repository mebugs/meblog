package platHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/platModel"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/plat/platService"
)

// AddAccount 	godoc
// @id			AddAccount 登陆账号新建
// @Summary		登陆账号新建
// @Description	新建登陆账号
// @Router		/plat/account/add [post]
// @Tags		登陆账号
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		platModel.AccountAddReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func AddAccount(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &platModel.AccountAddReq{})
	if err == nil {
		req := reqObj.(*platModel.AccountAddReq)
		// 执行创建
		service.JsonRes(c, platService.AddAccount(traceID, req))
	}
}

// PageAccount	godoc
// @id			PageAccount 登陆账号分页
// @Summary		登陆账号分页
// @Description	分页查询登陆账号
// @Router		/plat/account/page [post]
// @Tags		登陆账号
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		platModel.AccountPageReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=baseModel.PageRes{list=[]platModel.AccountPageRes}}	"响应成功"
func PageAccount(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &platModel.AccountPageReq{})
	if err == nil {
		req := reqObj.(*platModel.AccountPageReq)
		// 执行查询
		service.JsonRes(c, platService.PageAccount(traceID, req))
	}
}

// GetAccount	godoc
// @id			GetAccount 登陆账号详情
// @Summary		登陆账号详情
// @Description	查询登陆账号详情
// @Router		/plat/account/get [post]
// @Tags		登陆账号
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.IdReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=platModel.AccountGetRes}	"响应成功"
func GetAccount(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, platService.GetAccount(traceID, req))
	}
}

// EditAccount 	godoc
// @id			EditAccount 登陆账号编辑
// @Summary		登陆账号编辑
// @Description	基于数据ID编辑登陆账号
// @Router		/plat/account/edit [post]
// @Tags		登陆账号
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		platModel.AccountEditReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func EditAccount(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &platModel.AccountEditReq{})
	if err == nil {
		req := reqObj.(*platModel.AccountEditReq)
		// 执行编辑
		service.JsonRes(c, platService.EditAccount(traceID, req))
	}
}

// DelAccount	godoc
// @id			DelAccount 登陆账号移除
// @Summary		登陆账号移除
// @Description	登陆账号移除处理
// @Router		/plat/account/del [post]
// @Tags		登陆账号
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req		body		baseModel.IdReq	true			"请求"
// @Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
func DelAccount(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行移除
		service.JsonRes(c, platService.DelAccount(traceID, req))
	}
}

// ResetAccount	godoc
// @id			ResetAccount 登陆账号重置
// @Summary		登陆账号重置
// @Description	登陆账号中止密码
// @Router		/plat/account/reset [post]
// @Tags		登陆账号
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req		body		baseModel.IdReq	true			"请求"
// @Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
func ResetAccount(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行移除
		service.JsonRes(c, platService.ResetAccount(traceID, req))
	}
}
