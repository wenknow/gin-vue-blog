package main

import (
	"github.com/wenknow/gin-vue-blog/server/global"
	"github.com/wenknow/gin-vue-blog/server/initialize"
)

func main() {
	global.VP = initialize.Viper()    // 初始化Viper
	global.LOG = initialize.Zap()     // 初始化zap日志库
	global.DB = initialize.Gorm()     // gorm连接数据库
	global.REDIS = initialize.Redis() // go-redis连接

	initialize.RunWindowsServer()
}
