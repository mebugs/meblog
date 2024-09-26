package worker

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/cacheModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"siteol.com/smart/src/common/utils/security"
	"siteol.com/smart/src/config"
	"strings"
	"time"
)

// PrintUrl 打印URL
func PrintUrl(url, traceID string) {
	log.InfoTF(traceID, "Req Path Is : %s", url)
}

// SetLang 设置语言上下文
func SetLang(c *gin.Context) {
	lang := c.GetHeader(constant.ContextLang)
	if lang == "" || lang == "null" {
		lang = "zh-CN"
	}
	c.Set(constant.ContextLang, lang)
}

// SetRouter 设置路由上下文
func SetRouter(c *gin.Context, url, traceID string) (router *cacheModel.CacheRouter, ng bool) {
	router = cacheModel.CacheRouterNormal
	// 尝试读取缓存
	cacheGet := false
	resCacheMap, err := cacheModel.GetRouterCache(traceID)
	if err == nil {
		// 取得实际的配置
		routerGet, ok := resCacheMap[url]
		if !ok {
			log.WarnTF(traceID, "Get %s RouterCache Empty", url)
		} else {
			cacheGet = true
			router = routerGet
		}
	}
	c.Set(constant.ContextRouterC, router)
	// 缓存未取得且未启动Debug（测试环境数据库可以不配置路由）
	if !cacheGet && !config.JsonConfig.Server.Debug {
		// 非Debug直接退出
		log.ErrorTF(traceID, "Get %s RouterCache NG", url, err)
		ng = true
		return
	}
	return
}

// SetAuthUser 设置授权用户信息
func SetAuthUser(c *gin.Context, traceID string) (authUser *cacheModel.CacheAuth, token string, ng bool) {
	ng = true
	token = c.GetHeader(constant.HeaderToken)
	if token == "" || token == "null" {
		return
	}
	authUser, err := cacheModel.GetAuthCache(traceID, token)
	if err != nil {
		return
	}
	c.Set(constant.ContextAuthUser, authUser)
	ng = false
	return
}

// SetReq 处理请求 请求日志、入库等
func SetReq(c *gin.Context, router *cacheModel.CacheRouter, url, traceID string) (ng bool) {
	origReq, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.ErrorTF(traceID, "Req Read Fail . Err Is : %v", err)
		ng = true
		return
	}
	// COPY 请求给 Body
	runReq := io.NopCloser(bytes.NewBuffer(origReq))
	c.Request.Body = runReq
	// 处理日志打印
	printStr := "Do Not Print"
	// 请求字符为空
	// JSON序列化 以及安全处理
	printSafeStr := security.SafeJson(string(origReq), router.ReqSecure)
	// 打印日志，尝试脱敏
	if router.ReqPrint {
		printStr = printSafeStr
	}
	log.InfoTF(traceID, "Req Body Is : %s", printStr)
	// 日志入库
	if router.LogInDb {
		now := time.Now()
		logInDb := &platDB.RouterLog{
			Id:         0,
			AppName:    config.JsonConfig.Server.Name,
			AppNode:    config.SysNode,
			AppTraceId: traceID,
			ReqIp:      getClientIP(c),
			ReqUrl:     url,
			ReqBody:    printSafeStr,
			ReqAt:      &now,
			ResStatus:  0,
			ResBody:    "",
			ResTime:    nil,
		}
		// 传入入库日志对象
		c.Set(constant.ContextRouterI, logInDb)
	}
	return
}

// getClientIP 获取客户端IP
func getClientIP(c *gin.Context) string {
	req := c.Request
	// 尝试从 X-Forwarded-For 头部获取 IP 地址（如果使用了反向代理）
	forwarded := req.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// X-Forwarded-For 可能包含多个 IP 地址，通常由逗号分隔
		// 第一个通常是客户端的真实 IP 地址
		ips := strings.Split(forwarded, ",")
		clientIP := strings.TrimSpace(ips[0])
		return clientIP
	}
	// 如果 X-Forwarded-For 头部不存在，则使用 RemoteAddr
	return req.RemoteAddr
}
