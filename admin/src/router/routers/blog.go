package routers

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/router/middleware"
	"siteol.com/smart/src/service/blog/blogHandler"
)

// BlogRouter 博客路由
func BlogRouter(router *gin.Engine) {
	blogRouter := router.Group("/blog", middleware.CommMiddleWare) // 授权中间件
	{
		// 资源配置表相关
		sourceRouter := blogRouter.Group("/source")
		{
			sourceRouter.POST("/add", blogHandler.AddSource)
			sourceRouter.POST("/page", blogHandler.PageSource)
			sourceRouter.POST("/get", blogHandler.GetSource)
			sourceRouter.POST("/edit", blogHandler.EditSource)
			sourceRouter.POST("/del", blogHandler.DelSource)
		}

		// 文章分类相关
		categoryRouter := blogRouter.Group("/category")
		{
			categoryRouter.POST("/add", blogHandler.AddCategory)
			categoryRouter.POST("/page", blogHandler.PageCategory)
			categoryRouter.POST("/get", blogHandler.GetCategory)
			categoryRouter.POST("/edit", blogHandler.EditCategory)
			categoryRouter.POST("/del", blogHandler.DelCategory)
			categoryRouter.POST("/merge", blogHandler.MergeCategory)
			categoryRouter.POST("/list", blogHandler.ListCategory)
		}

		// 文章或页面相关
		postRouter := blogRouter.Group("/post")
		{
			postRouter.POST("/add", blogHandler.AddPost)
			postRouter.POST("/page", blogHandler.PagePost)
			postRouter.POST("/get", blogHandler.GetPost)
			postRouter.POST("/edit", blogHandler.EditPost)
			postRouter.POST("/del", blogHandler.DelPost)
		}

		// 文章引用标签相关
		postTagRouter := blogRouter.Group("/postTag")
		{
			postTagRouter.POST("/add", blogHandler.AddPostTag)
			postTagRouter.POST("/page", blogHandler.PagePostTag)
			postTagRouter.POST("/get", blogHandler.GetPostTag)
			postTagRouter.POST("/edit", blogHandler.EditPostTag)
			postTagRouter.POST("/del", blogHandler.DelPostTag)
		}

		// 文章评论相关
		postCommentRouter := blogRouter.Group("/postComment")
		{
			postCommentRouter.POST("/add", blogHandler.AddPostComment)
			postCommentRouter.POST("/page", blogHandler.PagePostComment)
			postCommentRouter.POST("/get", blogHandler.GetPostComment)
			postCommentRouter.POST("/edit", blogHandler.EditPostComment)
			postCommentRouter.POST("/del", blogHandler.DelPostComment)
		}

	}
}
