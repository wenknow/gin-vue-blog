package response

import (
	"github.com/gin-gonic/gin"
	"github.com/wenknow/gin-vue-blog/server/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

var ErrNoRecord = gorm.ErrRecordNotFound

const (
	LOGOUT  = 600
	ERROR   = 400
	SUCCESS = 200
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, msg, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMsg(msg string, err error, c *gin.Context) {
	if err != nil {
		global.LOG.Error(msg, zap.Any("err", err))
	}
	Result(ERROR, map[string]interface{}{}, msg, c)
}

func FailWithDetailed(data interface{}, msg string, c *gin.Context) {
	Result(ERROR, data, msg, c)
}

func FailWithLogout(msg string, c *gin.Context) {
	Result(LOGOUT, map[string]interface{}{}, msg, c)
}
