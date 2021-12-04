package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/wenknow/gin-vue-blog/server/global"
	"github.com/wenknow/gin-vue-blog/server/model/response"
	"github.com/wenknow/gin-vue-blog/server/utils/verify"
	"go.uber.org/zap"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 X-Token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("X-Token")
		if token == "" {
			response.FailWithLogout("未登录或非法访问", c)
			c.Abort()
			return
		}
		j := verify.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == verify.TokenExpired {
				response.FailWithLogout("登录信息过期", c)
				c.Abort()
				return
			}
			response.FailWithLogout("登录信息错误", c)
			global.LOG.Error("登录信息错误", zap.Any("err", err))
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
