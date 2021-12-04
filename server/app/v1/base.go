package v1

import (
	"github.com/wenknow/gin-vue-blog/server/app/v1/blog"
	"github.com/wenknow/gin-vue-blog/server/app/v1/system"
	"github.com/wenknow/gin-vue-blog/server/app/v1/user"
)

type ApiGroup struct {
	UserApiGroup    user.ApiGroup
	ArticleApiGroup blog.ApiGroup
	SystemApiGroup  system.ApiGroup
}

var NewApiGroup = new(ApiGroup)
