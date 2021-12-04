package blog

import (
	"github.com/wenknow/gin-vue-blog/server/service"
)

type ApiGroup struct {
	ArticleApi
	ArticleCgApi
	TagApi
	UserArticleGoodApi
	UserArticleCommentApi
}

var tagService = service.NewService.TagService
var userService = service.NewService.UserService
var articleService = service.NewService.ArticleService
var articleCgService = service.NewService.ArticleCgService
var articleTagService = service.NewService.ArticleTagService
var articleSearchService = service.NewService.ArticleSearchService
var userArticleGoodService = service.NewService.UserArticleGoodService
var userArticleBrowseService = service.NewService.UserArticleBrowseService
var userArticleCommentService = service.NewService.UserArticleCommentService
var userArticleCommentGoodService = service.NewService.UserArticleCommentGoodService
var defaultImgService = service.NewService.DefaultImgService
