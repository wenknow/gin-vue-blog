package repository

import "github.com/wenknow/gin-vue-blog/server/model/config"

type UserArticleGood struct {
	config.Model
	Uid       uint `json:"uid" form:"uid" gorm:"column:uid;comment:所有者;type:int"`
	ArticleId uint `json:"article_id" form:"article_id" gorm:"column:article_id;comment:文章id;type:int"`
}
