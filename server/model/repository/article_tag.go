package repository

import "github.com/wenknow/gin-vue-blog/server/model/config"

type ArticleTag struct {
	config.DelModel
	TagId     uint `json:"tagId" form:"tagId" gorm:"column:tag_id;comment:标签id;type:int"`
	Uid       uint `json:"uid" form:"uid" gorm:"column:uid;comment:所属者;type:int"`
	ArticleId uint `json:"article_id" form:"article_id" gorm:"column:article_id;comment:文章id;type:int"`
}
