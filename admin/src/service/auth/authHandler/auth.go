package authHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/auth/authService"
)

// AuthLogin 	godoc
// @id			AuthLogin 账号登陆（免授权）
// @Summary		账号密码登陆
// @Description	账号密码登陆
// @Router		/auth/login [post]
// @Tags		开放接口
// @Accept		json
// @Produce		json
// @Param		req	body		baseModel.AccountLoginReq		true				"请求"
// @Success		200	{object}	baseModel.ResBody{data=baseModel.AccountLoginRes}	"响应成功"
func AuthLogin(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.AccountLoginReq{})
	if err == nil {
		req := reqObj.(*baseModel.AccountLoginReq)
		// 执行创建
		service.JsonRes(c, authService.AuthLogin(traceID, req))
	}
}

// AuthLogout 	godoc
// @id			AuthLogout 账号登出（免授权）
// @Summary		账号登出
// @Description	账号登出
// @Router		/auth/logout [post]
// @Tags		开放接口
// @Accept		json
// @Produce		json
// @Param		req	body		baseModel.Req	true						"请求"
// @Success		200	{object}	baseModel.ResBody{data=baseModel.AccountLoginRes}	"响应成功"
func AuthLogout(c *gin.Context) {
	// traceID 日志追踪
	traceID := c.GetString(constant.ContextTraceID)
	// 尝试获取Token
	token := c.GetHeader(constant.HeaderToken)
	// Token下线异步处理
	if token != "" {
		go func() {
			authService.AuthLogout(traceID, token)
		}()
	}
	service.JsonRes(c, baseModel.SuccessUnPop(nil))
}

// AuthReset 	godoc
// @id			AuthReset 账号密码重置
// @Summary		账号密码重置
// @Description	账号密码重置
// @Router		/auth/reset [post]
// @Tags		开放接口
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.AccountResetReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}		"响应成功"
func AuthReset(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.AccountResetReq{})
	if err == nil {
		req := reqObj.(*baseModel.AccountResetReq)
		// 读取当前登陆用户
		authUser := service.GetAuthUser(c)
		// 执行创建
		service.JsonRes(c, authService.AuthReset(traceID, req, authUser))
	}
}

// AuthMine 	godoc
// @id			AuthMine 我的权限信息
// @Summary		我的权限信息
// @Description	我的权限信息
// @Router		/auth/mine [post]
// @Tags		开放接口
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.Req	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=baseModel.AuthMineRes}		"响应成功"
func AuthMine(c *gin.Context) {
	// 读取当前登陆用户
	authUser := service.GetAuthUser(c)
	if authUser != nil {
		service.JsonRes(c, baseModel.SuccessUnPop(baseModel.AuthMineRes{
			AccountId:      authUser.AccountId,
			Name:           authUser.Name,
			RoleNames:      authUser.RoleNames,
			PermissionKeys: authUser.PermissionKeys,
		}))
	} else {
		service.JsonRes(c, baseModel.ResFail)
	}
}
