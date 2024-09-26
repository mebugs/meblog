package middleware

import (
	"runtime/debug"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/service"

	"github.com/gin-gonic/gin"
)

// Recover 公共Panic
func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.ErrorF("panic:%s, stack:%s", err, string(debug.Stack()))
			service.JsonRes(c, baseModel.SysErr)
			c.Abort()
		}
	}()
	c.Next()
}
