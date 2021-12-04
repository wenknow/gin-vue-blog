package service

import (
	"github.com/wenknow/gin-vue-blog/server/global"
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"github.com/wenknow/gin-vue-blog/server/model/request"
	"github.com/wenknow/gin-vue-blog/server/model/response"
	"github.com/wenknow/gin-vue-blog/server/utils/upload"
	"mime/multipart"
	"strings"
)

type FileUploadService struct {
}

func (u *FileUploadService) GetFileUpload(cond map[string]interface{}) (FileUpload repository.FileUpload, err error) {
	err = buildCond(cond).First(&FileUpload).Error
	return
}

func (u *FileUploadService) GetFileUploadList(cond map[string]interface{}) (FileUploadList []repository.FileUpload, err error) {
	err = buildCond(cond).Find(&FileUploadList).Error
	return
}

func (u *FileUploadService) GetFileUploadByPk(id uint) (FileUpload repository.FileUpload, err error) {
	err = global.DB.First(&FileUpload, id).Error
	return
}

func (u *FileUploadService) PagerFileUpload(cond map[string]interface{}, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNum - 1)
	db := buildCond(cond).Model(&repository.FileUpload{})
	var FileUploadList []repository.FileUpload
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&FileUploadList).Error
	return FileUploadList, total, err
}

func (u *FileUploadService) AddFileUpload(FileUpload repository.FileUpload) (uid uint, err error) {
	err = global.DB.Create(&FileUpload).Error
	return FileUpload.ID, err
}

func (u *FileUploadService) SaveFileUpload(FileUpload repository.FileUpload) (err error) {
	err = global.DB.Save(&FileUpload).Error
	return
}

func (u *FileUploadService) UpdateFileUpload(FileUpload repository.FileUpload) (err error) {
	err = global.DB.Updates(&FileUpload).Error
	return
}

func (u *FileUploadService) DelFileUpload(id uint) (err error) {
	err = global.DB.Delete(&repository.FileUpload{}, id).Error
	return
}

func (u *FileUploadService) UploadFile(header *multipart.FileHeader, uid uint) (file repository.FileUpload, err error) {

	//同一个用户同一个文件名 相同大小的文件不上传
	cond := InitCond()
	cond["uid"] = uid
	cond["name"] = header.Filename
	cond["size"] = header.Size
	if oldFile, err := u.GetFileUpload(cond); err == response.ErrNoRecord {
	} else if err != nil {
		return file, err
	} else {
		return oldFile, nil
	}

	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	s := strings.Split(header.Filename, ".")
	f := repository.FileUpload{
		Uid:  uid,
		Url:  filePath,
		Name: header.Filename,
		Tag:  s[len(s)-1],
		Key:  key,
		Size: uint(header.Size),
	}
	newId, err := u.AddFileUpload(f)
	f.ID = newId
	return f, err
}
