package middleware

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/cacheModel"
	"siteol.com/smart/src/common/utils"
	"siteol.com/smart/src/router/middleware/worker"
	"siteol.com/smart/src/service"
)

// CommMiddleWare 开放中间件
func CommMiddleWare(c *gin.Context) {
	// 生成请求唯一标志
	traceID := utils.TraceID()
	url := c.Request.URL.Path
	worker.PrintUrl(url, traceID)
	c.Set(constant.ContextTraceID, traceID)
	// 设置语言
	worker.SetLang(c)
	// 读取路由配置，可能会NG 404 路由找不到
	router, routerNg := worker.SetRouter(c, url, traceID)
	if routerNg {
		c.Abort()
		service.JsonRes(c, baseModel.PathErr)
		return
	}
	// 请求日志处理，可能会NG 500 系统异常
	reqNg := worker.SetReq(c, router, url, traceID)
	if reqNg {
		c.Abort()
		service.JsonRes(c, baseModel.SysErr)
		return
	}
	// 读取鉴权信息
	if router.NeedAuth {
		authUser, token, ng := worker.SetAuthUser(c, traceID)
		if ng {
			c.Abort()
			service.JsonRes(c, baseModel.LoginErr) // 尚未登陆或登录过期
			return
		}
		if !utils.ArrayInclude(url, authUser.Routers) {
			c.Abort()
			service.JsonRes(c, baseModel.AuthErr) // 无权访问此接口
			return
		}
		// 授权时间续杯
		err := cacheModel.RefreshAuthCache(traceID, token)
		if err != nil {
			c.Abort()
			service.JsonRes(c, baseModel.SysErr)
			return
		}
	}
	c.Next()
}
