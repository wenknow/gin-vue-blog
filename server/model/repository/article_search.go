package repository

import "github.com/wenknow/gin-vue-blog/server/model/config"

type ArticleSearch struct {
	config.DelModel
	Content string `json:"content" form:"content" gorm:"column:content;comment:内容;type:text;"`
}
