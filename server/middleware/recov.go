package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Api struct {
	Code    int
	Message string
}

// Handler 异常处理
func Handler(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			switch t := r.(type) {
			case *Api:
				log.Printf("panic: %v\n", t.Message)
				c.JSON(t.Code, gin.H{
					"message": t.Message,
				})
			default:
				log.Printf("panic: internal error")
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "服务器内部异常",
				})
			}
			c.Abort()
		}
	}()

	c.Next()
}
