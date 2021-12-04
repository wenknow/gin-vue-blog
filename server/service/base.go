package service

import (
	"github.com/wenknow/gin-vue-blog/server/global"
	"github.com/wenknow/gin-vue-blog/server/model/request"
	"gorm.io/gorm"
)

type Service struct {
	UserService
	ApiLogService
	ArticleService
	ArticleCgService
	ArticleSearchService
	TagService
	ArticleTagService
	FileUploadService
	UserArticleGoodService
	UserArticleBrowseService
	UserArticleCommentService
	UserArticleCommentGoodService
	DefaultImgService
}

var NewService = new(Service)

func InitCond() map[string]interface{} {
	cond := make(map[string]interface{})
	return cond
}

func InitPageInfo(pageNum int, pageSize int) (info request.PageInfo) {
	if pageNum == 0 {
		pageNum = 1
	}
	info.PageNum = pageNum
	if pageSize == 0 {
		pageSize = global.CONFIG.System.PageSize
	}
	info.PageSize = pageSize
	return
}

func buildCond(cond map[string]interface{}) *gorm.DB {
	db := global.DB
	for k, v := range cond {
		if k == "OR" {
			db = db.Or(v)
		} else if k == "NOT" {
			db = db.Not(v)
		} else if k == "ORDER" {
			db = db.Order(v)
		} else if k == "LIMIT" {
			db = db.Limit(v.(int))
		} else {
			db = db.Where(k, v)
		}
	}
	return db
}
