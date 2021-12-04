package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wenknow/gin-vue-blog/server/model/response"
	"github.com/wenknow/gin-vue-blog/server/utils/verify"
	"path/filepath"
)

type FileUploadApi struct {
}

func (u *FileUploadApi) UploadImg(c *gin.Context) {
	header, err := c.FormFile("image")
	if err != nil {
		response.FailWithMsg("接收文件失败!", err, c)
		return
	}
	userInfo := verify.GetUserInfo(c)

	if header.Size > 1024*1024*1 {
		response.FailWithMsg("图片大小不能超过1M", err, c)
		return
	}

	fileExt := filepath.Ext(header.Filename)
	allowExts := []string{".jpg", ".png", ".gif", ".jpeg"}
	allowFlag := false
	for _, ext := range allowExts {
		if ext == fileExt {
			allowFlag = true
			break
		}
	}
	if !allowFlag {
		response.FailWithMsg("不允许的文件类型", err, c)
		return
	}
	file, err := fileUploadService.UploadFile(header, userInfo.ID)
	if err != nil {
		response.FailWithMsg("上传失败", err, c)
		return
	}

	response.OkWithData(file, c)
}
