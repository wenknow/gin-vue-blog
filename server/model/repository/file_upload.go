package repository

import "github.com/wenknow/gin-vue-blog/server/model/config"

type FileUpload struct {
	config.Model
	Uid  uint   `json:"uid" gorm:"comment:所属用户;type:int"` // 所属用户
	Name string `json:"name" gorm:"comment:文件名"`          // 文件名
	Url  string `json:"url" gorm:"comment:文件地址"`          // 文件地址
	Tag  string `json:"tag" gorm:"comment:文件标签"`          // 文件标签
	Key  string `json:"key" gorm:"comment:编号"`            // 编号
	Size uint   `json:"size" gorm:"comment:大小;type:int"`  // 大小
}
