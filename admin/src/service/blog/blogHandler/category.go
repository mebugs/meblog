package blogHandler

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/blogModel"
	"siteol.com/smart/src/service"
	"siteol.com/smart/src/service/blog/blogService"
)

// AddCategory 	godoc
// @id			AddCategory 文章分类新建
// @Summary		文章分类新建
// @Description	新建文章分类
// @Router		/blog/category/add [post]
// @Tags		文章分类
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.CategoryAddReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func AddCategory(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.CategoryAddReq{})
	if err == nil {
		req := reqObj.(*blogModel.CategoryAddReq)
		// 执行创建
		service.JsonRes(c, blogService.AddCategory(traceID, req))
	}
}

// PageCategory	godoc
// @id			PageCategory 文章分类分页
// @Summary		文章分类分页
// @Description	分页查询文章分类
// @Router		/blog/category/page [post]
// @Tags		文章分类
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.CategoryPageReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=baseModel.PageRes{list=[]blogModel.CategoryPageRes}}	"响应成功"
func PageCategory(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.CategoryPageReq{})
	if err == nil {
		req := reqObj.(*blogModel.CategoryPageReq)
		// 执行查询
		service.JsonRes(c, blogService.PageCategory(traceID, req))
	}
}

// GetCategory	godoc
// @id			GetCategory 文章分类详情
// @Summary		文章分类详情
// @Description	查询文章分类详情
// @Router		/blog/category/get [post]
// @Tags		文章分类
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		baseModel.IdReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=blogModel.CategoryGetRes}	"响应成功"
func GetCategory(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		// 执行查询
		service.JsonRes(c, blogService.GetCategory(traceID, req))
	}
}

// EditCategory 	godoc
// @id			EditCategory 文章分类编辑
// @Summary		文章分类编辑
// @Description	基于数据ID编辑文章分类
// @Router		/blog/category/edit [post]
// @Tags		文章分类
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req	body		blogModel.CategoryEditReq	true	"请求"
// @Success		200	{object}	baseModel.ResBody{data=bool}	"响应成功"
func EditCategory(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.CategoryEditReq{})
	if err == nil {
		req := reqObj.(*blogModel.CategoryEditReq)
		// 执行编辑
		service.JsonRes(c, blogService.EditCategory(traceID, req))
	}
}

// DelCategory	godoc
// @id			DelCategory 文章分类移除
// @Summary		文章分类移除
// @Description	文章分类移除处理
// @Router		/blog/category/del [post]
// @Tags		文章分类
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req		body		baseModel.IdReq	true			"请求"
// @Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
func DelCategory(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &baseModel.IdReq{})
	if err == nil {
		req := reqObj.(*baseModel.IdReq)
		service.JsonRes(c, blogService.DelCategory(traceID, req))
	}
}

// MergeCategory	godoc
// @id			MergeCategory 文章分类合并
// @Summary		文章分类合并
// @Description	文章分类合并处理
// @Router		/blog/category/merge [post]
// @Tags		文章分类
// @Accept		json
// @Produce		json
// @Security	Token
// @Param		req		body		blogModel.CategoryToReq	true			"请求"
// @Success		200		{object}	baseModel.ResBody{data=bool}	"响应成功"
func MergeCategory(c *gin.Context) {
	traceID, reqObj, err := service.ValidateReqObj(c, &blogModel.CategoryToReq{})
	if err == nil {
		req := reqObj.(*blogModel.CategoryToReq)

		service.JsonRes(c, blogService.MergeCategory(traceID, req))
	}
}

// ListCategory	godoc
// @id			ListCategory 文章分类列表
// @Summary		文章分类列表
// @Description	文章分类列表处理
// @Router		/blog/category/list [post]
// @Tags		文章分类
// @Accept		json
// @Produce		json
// @Security	Token
// @Success		200		{object}	baseModel.ResBody{data=[]baseModel.SelectNumRes}	"响应成功"
func ListCategory(c *gin.Context) {
	traceID := c.GetString(constant.ContextTraceID)
	service.JsonRes(c, blogService.ListCategory(traceID))
}
