package repository

import "github.com/wenknow/gin-vue-blog/server/model/config"

type DefaultImg struct {
	config.Model
	Uid    uint   `json:"uid" gorm:"comment:所属用户;type:int"`  // 所属用户
	Name   string `json:"name" gorm:"comment:文件名"`           // 文件名
	Url    string `json:"url" gorm:"comment:文件地址"`           // 文件地址
	Type   string `json:"type" gorm:"comment:文件类型"`          // 文件类型
	Height uint   `json:"height" gorm:"comment:高度;type:int"` // 高度
	Weight uint   `json:"weight" gorm:"comment:宽度;type:int"` // 宽度
	Size   uint   `json:"size" gorm:"comment:大小;type:int"`   // 大小
}
