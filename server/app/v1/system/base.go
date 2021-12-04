package system

import "github.com/wenknow/gin-vue-blog/server/service"

type ApiGroup struct {
	FileUploadApi
}

var fileUploadService = service.NewService.FileUploadService
