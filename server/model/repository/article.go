package repository

import "github.com/wenknow/gin-vue-blog/server/model/config"

type Article struct {
	config.Model
	BrowseCount  uint   `json:"browse_count" form:"browse_count" gorm:"column:browse_count;comment:浏览量;type:int"`
	CgId         uint   `json:"cg_id" form:"cg_id" gorm:"column:cg_id;comment:分类id;type:int"`
	CollectCount uint   `json:"collect_count" form:"collect_count" gorm:"column:collect_count;comment:收藏量;type:int"`
	CommentCount uint   `json:"comment_count" form:"comment_count" gorm:"column:comment_count;comment:评论量;type:int"`
	WordCount    uint   `json:"word_count" form:"word_count" gorm:"column:word_count;comment:字数;type:int"`
	Content      string `json:"content" form:"content" gorm:"column:content;comment:内容;type:text;"`
	ContentHtml  string `json:"content_html" form:"content_html" gorm:"column:content_html;comment:内容;type:text;"`
	GoodCount    uint   `json:"good_count" form:"good_count" gorm:"column:good_count;comment:点赞量;type:int"`
	PublicIs     bool   `json:"public_is" form:"public_is" gorm:"column:public_is;comment:是否公开;type:tinyint"`
	Title        string `json:"title" form:"title" gorm:"column:title;comment:标题;type:varchar(80);"`
	Desc         string `json:"desc" form:"desc" gorm:"column:desc;comment:简介;type:varchar(255);"`
	Tags         string `json:"tags" form:"tags" gorm:"column:tags;comment:标签名;type:varchar(255);"`
	ImgUrl       string `json:"img_url" form:"img_url" gorm:"column:img_url;comment:封面图;type:varchar(255);"`
	TopIs        bool   `json:"top_is" form:"top_is" gorm:"column:top_is;comment:是否置顶;type:tinyint"`
	Uid          uint   `json:"uid" form:"uid" gorm:"column:uid;comment:所属者;type:int"`
}
