package repository

import "github.com/wenknow/gin-vue-blog/server/model/config"

type ArticleCg struct {
	config.Model
	Name string `json:"name" form:"name" gorm:"column:name;comment:名称;type:varchar(30);"`
	Sort int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序值;type:int"`
	Uid  uint   `json:"uid" form:"uid" gorm:"column:uid;comment:所属者;type:int"`
}
