package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/wenknow/gin-vue-blog/server/app/v1"
	"github.com/wenknow/gin-vue-blog/server/middleware"
)

func (a *ApiRouter) ArticleRouter(Router *gin.RouterGroup) {

	var tagApi = v1.NewApiGroup.ArticleApiGroup.TagApi
	var articleApi = v1.NewApiGroup.ArticleApiGroup.ArticleApi
	var articleCgApi = v1.NewApiGroup.ArticleApiGroup.ArticleCgApi
	var userArticleGoodApi = v1.NewApiGroup.ArticleApiGroup.UserArticleGoodApi
	var userArticleCommentApi = v1.NewApiGroup.ArticleApiGroup.UserArticleCommentApi

	apiRouter := Router.Group("/blog").Use(middleware.JWTAuth()).Use(middleware.OperationRecord())
	apiRouterNoLog := Router.Group("/blog").Use(middleware.JWTAuth())
	apiRouterRaw := Router.Group("/blog")
	{
		apiRouterRaw.GET("info", articleApi.GetInfo) // 新增分类

		apiRouterRaw.POST("list", articleApi.GetArticleList)               // 新增分类
		apiRouterRaw.POST("detail", articleApi.GetArticleDetail)           // 新增分类
		apiRouterRaw.POST("userList", articleApi.GetUserArticleList)       // 获取用户文章列表
		apiRouterRaw.POST("relatedList", articleApi.GetRelatedArticleList) // 获取相关文章

		apiRouterNoLog.POST("save", articleApi.SaveArticle)  // 自动保存
		apiRouter.POST("publish", articleApi.PublishArticle) // 发布文章
		apiRouter.POST("del", articleApi.DelArticle)         // 删除分类

		apiRouterRaw.GET("generateTags", articleApi.GenerateTags) // 脚本生成tags
	}

	cgApiRouterRaw := Router.Group("/blog/cg")
	{
		cgApiRouterRaw.POST("add", articleCgApi.AddArticleCg)       // 新增分类
		cgApiRouterRaw.POST("update", articleCgApi.UpdateArticleCg) // 更新分类
		cgApiRouterRaw.POST("del", articleCgApi.DelArticleCg)       // 删除分类
		cgApiRouterRaw.POST("list", articleCgApi.GetArticleCgList)  // 获取分类列表
	}

	tagApiRouterRaw := Router.Group("/blog/tag")
	{
		tagApiRouterRaw.POST("guess", tagApi.GuessTag)          // 根据文章内容推测标签
		tagApiRouterRaw.POST("userList", tagApi.GetUserTagList) // 获取用户标签列表
		tagApiRouterRaw.POST("list", tagApi.GetTagList)         // 获取全部标签列表
		tagApiRouterRaw.POST("main", tagApi.GetMainTagList)     // 获取已使用标签列表
	}

	goodApiRouterRaw := Router.Group("/blog/good")
	{
		goodApiRouterRaw.POST("add", userArticleGoodApi.AddUserArticleGood) // 文章点赞
		goodApiRouterRaw.POST("del", userArticleGoodApi.DelUserArticleGood) // 文章取消点赞
	}

	commentApiRouterRaw := Router.Group("/blog/comment").Use(middleware.JWTAuth()).Use(middleware.OperationRecord())
	commentApiRouterRawNoLog := Router.Group("/blog/comment")
	{
		commentApiRouterRaw.POST("add", userArticleCommentApi.AddUserArticleComment)           // 文章点赞
		commentApiRouterRaw.POST("reply", userArticleCommentApi.AddUserArticleReply)           // 文章取消点赞
		commentApiRouterRaw.POST("del", userArticleCommentApi.DelUserArticleComment)           // 删除文章评论
		commentApiRouterRawNoLog.POST("list", userArticleCommentApi.GetUserArticleCommentList) // 获取文章评论列表
	}

}
