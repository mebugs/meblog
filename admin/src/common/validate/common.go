package validate

import (
	"errors"
	"net/http"
	"siteol.com/smart/src/common/model/baseModel"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
)

// Readable 请求数据解析和校验错误返回
func Readable(c *gin.Context, req any) (*baseModel.ResBody, any) {
	traceID := c.GetString(constant.ContextTraceID)
	var err error
	// 根据请求采用不同绑定
	if c.Request.Method == http.MethodGet {
		err = c.ShouldBindQuery(req)
	} else if c.Request.Method == http.MethodPost {
		err = c.ShouldBindJSON(req)
	}
	// 校验出错
	if err != nil {
		errResult := readableError(err, c.GetString(constant.ContextLang), traceID)
		// 响应400错误（已翻译）
		return baseModel.Validate(errResult), nil
	}
	return nil, req
}

// 解析异常，如果是校验异常，返回自定义结果
func readableError(err error, lang, traceID string) error {
	// 翻译校验错误
	if errs, ok := err.(validator.ValidationErrors); ok {
		log.DebugTF(traceID, "ValidationErrors : %v", errs)
		return transErr(errs, lang)
	}
	// 未知校验错误
	log.ErrorTF(traceID, "ReadableError fail error: %s", err)
	return errors.New("UnknownValidationError")
}

// 翻译校验错误
func transErr(errs validator.ValidationErrors, lang string) error {
	// 优先取用第一个校验异常
	err := errs[0]
	// fieldName := err.Namespace() // 格式为：AuthoriseReq.StoreInfo.StoreNum
	// 翻译错误校验国际化
	return errors.New(err.Translate(transMap[lang]))
}
