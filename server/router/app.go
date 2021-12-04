package router

import (
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
}

func AppRouter() *gin.Engine {

	router := gin.Default()
	apiRouter := new(ApiRouter)

	//基础路由
	baseGroup := router.Group("/api")

	{
		apiRouter.UserRouter(baseGroup)
		apiRouter.ArticleRouter(baseGroup)
		apiRouter.SystemRouter(baseGroup)
	}
	return router
}
