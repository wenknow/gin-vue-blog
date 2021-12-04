package repository

import "github.com/wenknow/gin-vue-blog/server/model/config"

type UserArticleComment struct {
	config.Model
	Uid            uint   `json:"uid" form:"uid" gorm:"column:uid;comment:所有者;type:int"`
	ToUid          uint   `json:"to_uid" form:"to_uid" gorm:"column:to_uid;comment:被回复的用户;type:int"`
	ArticleId      uint   `json:"article_id" form:"article_id" gorm:"column:article_id;comment:文章id;type:int"`
	Content        string `json:"content" form:"content" gorm:"column:content;comment:内容;type:text;"`
	ToCommentId    uint   `json:"to_comment_id" form:"to_comment_id" gorm:"column:to_comment_id;comment:被回复的评论;type:int"`
	ToSupCommentId uint   `json:"to_sup_comment_id" form:"to_sup_comment_id" gorm:"column:to_sup_comment_id;comment:被回复的主评论id;type:int"`
	GoodCount      uint   `json:"good_count" form:"good_count" gorm:"column:good_count;comment:点赞数;type:int"`
	ReplayCount    uint   `json:"replay_count" form:"replay_count" gorm:"column:replay_count;comment:回复数;type:int"`
	Ip             string `json:"ip" form:"ip" gorm:"column:ip;comment:ip;type:varchar(40)"`
	Address        string `json:"address" form:"address" gorm:"column:address;comment:地址;type:varchar(60)"`
}
