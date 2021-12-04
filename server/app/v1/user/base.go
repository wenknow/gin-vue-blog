package user

import (
	"context"
	"github.com/wenknow/gin-vue-blog/server/service"
)

type ApiGroup struct {
	LoginApi
	UserApi
}

var userService = service.NewService.UserService
var articleService = service.NewService.ArticleService
var ctxB = context.Background()
