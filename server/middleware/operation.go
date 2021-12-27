package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/wenknow/gin-vue-blog/server/global"
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"github.com/wenknow/gin-vue-blog/server/service"
	"github.com/wenknow/gin-vue-blog/server/utils/verify"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

var apiLogService = service.NewService.ApiLogService

func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		var err error

		//get请求不做记录
		if c.Request.Method == http.MethodGet {
			return
		}

		body, err = ioutil.ReadAll(c.Request.Body)
		if err != nil {
			global.LOG.Error("读取请求出错:", zap.Any("err", err))
		} else {
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}

		myUid := verify.GetTokenUid(c)
		record := repository.ApiLog{
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Agent:  c.Request.UserAgent(),
			Body:   string(body),
			Uid:    myUid,
		}

		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		latency := time.Since(now).String()
		record.Error = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Latency = latency
		record.Resp = writer.body.String()

		if _, err := apiLogService.AddApiLog(record); err != nil {
			global.LOG.Error("新增api日志出错:", zap.Any("err", err))
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
