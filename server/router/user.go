package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/wenknow/gin-vue-blog/server/app/v1"
	"github.com/wenknow/gin-vue-blog/server/middleware"
)

func (a *ApiRouter) UserRouter(Router *gin.RouterGroup) {

	var loginApi = v1.NewApiGroup.UserApiGroup.LoginApi
	var userApi = v1.NewApiGroup.UserApiGroup.UserApi
	apiRouter := Router.Group("/user").Use(middleware.OperationRecord())
	apiRouterLogin := Router.Group("/user").Use(middleware.JWTAuth()).Use(middleware.OperationRecord())
	{
		apiRouter.POST("login", loginApi.Login)                 // 用户登录
		apiRouter.POST("logout", loginApi.Logout)               // 退出登录
		apiRouter.POST("sendEmailCode", loginApi.SendEmailCode) // 发送邮箱验证码
		apiRouter.POST("author", userApi.GetAuthorInfo)         // 获取作者信息
		apiRouterLogin.POST("info", userApi.GetUserInfo)        // 获取用户信息
	}

}
