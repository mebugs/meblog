package baseHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/service"
	"strings"
)

// Sample 		godoc
// @id			Sample示例
// @Summary		通用API示例
// @Description	系统API基本示例
// @Router		/docs/sample [post]
// @Tags		开放接口
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.DocsReq	true	"示例请求"
// @Success		200	{object}	baseModel.DocsOk			"业务受理成功"
// @Failure		400	{object}	baseModel.DocsVail			"数据校验失败"
// @Failure		401	{object}	baseModel.DocsAuthLg		"当前尚未登陆"
// @Failure		403	{object}	baseModel.DocsAuthNg		"权限校验失败"
// @Failure		500	{object}	baseModel.DocsErr			"服务系统异常"
func Sample(c *gin.Context) {
	_, req, err := service.ValidateReqObj(c, &baseModel.DocsReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, baseModel.DocsVail{Code: constant.ValidErr, Msg: "参数非法"})
		return
	}
	DocsReq := req.(*baseModel.DocsReq)
	switch DocsReq.HttpCode {
	case http.StatusInternalServerError: // 500
		c.JSON(http.StatusInternalServerError, baseModel.DocsErr{Code: constant.Error, Msg: "系统异常"})
	case http.StatusBadRequest: // 400
		c.JSON(http.StatusBadRequest, baseModel.DocsVail{Code: constant.ValidErr, Msg: "参数非法"})
	case http.StatusUnauthorized: // 401
		c.JSON(http.StatusBadRequest, baseModel.DocsAuthLg{Code: constant.LoginErr, Msg: "当前尚未登陆"})
	case http.StatusForbidden: // 403
		c.JSON(http.StatusBadRequest, baseModel.DocsAuthNg{Code: constant.AuthErr, Msg: "禁止访问"})
	default:
		c.JSON(http.StatusOK, baseModel.DocsOk{Code: constant.Success, Msg: "业务请求成功"})
	}
	return
}

// ReDoc HTML加载
func ReDoc(c *gin.Context) {
	c.Data(http.StatusOK, "text/html", []byte(constant.DocHtml))
	return
}

// SwaggerDoc HTML加载
func SwaggerDoc(c *gin.Context) {
	c.Data(http.StatusOK, "text/html", []byte(constant.SwaggerHtml))
	return
}

// DocsFile 静态文件
func DocsFile(c *gin.Context) {
	url := c.Request.URL.Path
	fileInfo := url[strings.LastIndex(url, "/"):]
	fileEnd := fileInfo[strings.LastIndex(fileInfo, "."):]
	contextType := "application/octet-stream"
	switch fileEnd {
	case ".png":
		contextType = "image/png"
	case ".js":
		contextType = "application/javascript"
	case ".css":
		contextType = "text/css"
	}
	data, err := os.ReadFile("docs/" + fileInfo)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
	} else {
		c.Data(http.StatusOK, contextType, data)
	}
	return
}
