package repository

import "github.com/wenknow/gin-vue-blog/server/model/config"

type UserArticleCommentGood struct {
	config.Model
	Uid       uint `json:"uid" form:"uid" gorm:"column:uid;comment:所有者;type:int"`
	ArticleId uint `json:"article_id" form:"article_id" gorm:"column:article_id;comment:文章id;type:int"`
	CommentId uint `json:"comment_id" form:"comment_id" gorm:"column:comment_id;comment:文章id;type:int"`
}
