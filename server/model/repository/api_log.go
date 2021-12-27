package repository

import (
	"github.com/wenknow/gin-vue-blog/server/model/config"
)

// 如果含有time.Time 请自行import time包
type ApiLog struct {
	config.Model
	Ip      string `json:"ip" form:"ip" gorm:"column:ip;comment:请求ip;type:varchar(40)"`                  // 请求ip
	Method  string `json:"method" form:"method" gorm:"column:method;comment:请求方法;type:varchar(20)"`      // 请求方法
	Path    string `json:"path" form:"path" gorm:"column:path;comment:请求路径"`                             // 请求路径
	Status  int    `json:"status" form:"status" gorm:"column:status;comment:请求状态"`                       // 请求状态
	Latency string `json:"latency" form:"latency" gorm:"column:latency;comment:延迟" swaggertype:"string"` // 延迟
	Agent   string `json:"agent" form:"agent" gorm:"column:agent;comment:代理"`                            // 代理
	Error   string `json:"error" form:"error" gorm:"column:error;comment:错误信息"`                          // 错误信息
	Body    string `json:"body" form:"body" gorm:"type:text;column:body;comment:请求Body"`                 // 请求Body
	Resp    string `json:"resp" form:"resp" gorm:"type:text;column:resp;comment:响应Body"`                 // 响应Body
	Uid     uint   `json:"uid" form:"uid" gorm:"column:uid;comment:用户id"`                                // 用户id
}
