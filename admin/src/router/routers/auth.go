package routers

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/router/middleware"
	"siteol.com/smart/src/service/auth/authHandler"
)

// AuthRouter 授权路由
func AuthRouter(router *gin.Engine) {
	// 授权域名在路由配置中定义为白名单路由，所有默认无需授权，因此可以共用中间件
	authRouter := router.Group("/auth", middleware.CommMiddleWare)
	{
		authRouter.POST("/login", authHandler.AuthLogin)   // 账密登陆（免授权）
		authRouter.POST("/logout", authHandler.AuthLogout) // 账号登出（免授权）
		authRouter.POST("/reset", authHandler.AuthReset)   // 密码重置
		authRouter.POST("/mine", authHandler.AuthMine)     // 我的权限信息
	}
}
