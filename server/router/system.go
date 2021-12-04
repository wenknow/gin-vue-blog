package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/wenknow/gin-vue-blog/server/app/v1"
	"github.com/wenknow/gin-vue-blog/server/middleware"
)

func (a *ApiRouter) SystemRouter(Router *gin.RouterGroup) {

	var fileUploadApi = v1.NewApiGroup.SystemApiGroup.FileUploadApi
	apiRouter := Router.Group("/system").Use(middleware.JWTAuth())
	{
		apiRouter.POST("uploadImg", fileUploadApi.UploadImg) // 文件上传
	}

}
