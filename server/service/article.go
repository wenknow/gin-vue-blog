package service

import "C"
import (
	"github.com/wenknow/gin-vue-blog/server/global"
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"github.com/wenknow/gin-vue-blog/server/model/request"
	"github.com/wenknow/gin-vue-blog/server/model/response"
	"gorm.io/gorm"
)

type ArticleService struct {
}

func (u *ArticleService) GetAuthorInfo(uid uint) (articleInfo response.ArticleGroupByUserInfo, err error) {
	err = global.DB.Table("user A").Select("A.id, A.name, A.head_url, A.desc,A.gender,"+
		" count(B.id) as all_article_count, sum(B.browse_count) as all_browse_count, "+
		"sum(B.good_count) as all_good_count, sum(B.collect_count) as all_collect_count, sum(B.comment_count) as all_comment_count").
		Joins("left join article B on B.uid = A.id AND B.public_is = true").
		Where("A.id", uid).Group("A.id").Take(&articleInfo).Error
	return
}

func (u *ArticleService) GetArticleGroupByUid(uid uint) (articleInfo response.ArticleGroupByUserInfo, err error) {
	err = global.DB.Table("article A").Select("B.id, B.name, B.head_url, B.desc,B.gender,"+
		" count(A.id) as all_article_count, sum(A.browse_count) as all_browse_count, "+
		"sum(A.good_count) as all_good_count, sum(A.collect_count) as all_collect_count, sum(A.comment_count) as all_comment_count").
		Joins("left join user B on A.uid = B.id").
		Where("uid", uid).Where("A.public_is", true).Group("B.id").Take(&articleInfo).Error
	return
}

func (u *ArticleService) GetArticleInfoByPk(id uint) (articleInfo response.ArticleInfo, err error) {

	err = global.DB.Table("article").Select("article.*, article_cg.name as cg_name").
		Joins("left join article_cg ON article.cg_id = article_cg.id").Where("article.id", id).Take(&articleInfo).Error
	return
}

func (u *ArticleService) GetArticle(cond map[string]interface{}) (article repository.Article, err error) {
	err = buildCond(cond).First(&article).Error
	return
}

func (u *ArticleService) GetArticleList(cond map[string]interface{}) (articleCgList []repository.Article, err error) {
	err = buildCond(cond).Find(&articleCgList).Error
	return
}

func (u *ArticleService) GetArticleByPk(id uint) (article repository.Article, err error) {
	err = global.DB.First(&article, id).Error
	return
}

func (u *ArticleService) PagerArticle(cond map[string]interface{}, info request.PageInfo) (list []response.ArticleListInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNum - 1)
	db := buildCond(cond).Model(&repository.Article{}).Joins("LEFT JOIN user B on article.uid = B.id").Joins("LEFT JOIN article_cg C on article.cg_id = C.id")
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Select("article.id, article.title, article.desc, article.uid, article.img_url, article.top_is, article.word_count, " +
		"article.browse_count, article.good_count,  article.tags,article.collect_count, article.comment_count, article.created_at," +
		"article.updated_at, B.name as author_name, B.head_url as author_head_url, C.name as cg_name").
		Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err
}

func (u *ArticleService) AddArticle(article repository.Article) (id uint, err error) {
	err = global.DB.Create(&article).Error
	return article.ID, err
}

func (u *ArticleService) SaveArticle(article repository.Article) (err error) {
	err = global.DB.Save(&article).Error
	return
}

func (u *ArticleService) UpdateArticle(article repository.Article) (err error) {
	err = global.DB.Updates(&article).Error
	return
}

func (u *ArticleService) DelArticle(id uint) (err error) {
	err = global.DB.Delete(&repository.Article{}, id).Error
	return
}

func (u *ArticleService) GetArticleByNameUid(name string, uid uint) (article repository.Article, err error) {
	err = global.DB.Where("name", name).Where("uid", uid).First(&article).Error
	return
}

func (u *ArticleService) GetArticleListByUid(uid uint) (articleCgList []repository.Article, err error) {
	err = global.DB.Where("uid", uid).Find(&articleCgList).Error
	return
}

func (u *ArticleService) UpdateArticleBrowseCount(id uint) (err error) {
	err = global.DB.Model(&repository.Article{}).Where("id", id).
		UpdateColumn("browse_count", gorm.Expr("browse_count+ ?", 1)).Error
	return
}
func (u *ArticleService) AddArticleCommentCount(id uint) (err error) {
	err = global.DB.Model(&repository.Article{}).Where("id", id).
		UpdateColumn("comment_count", gorm.Expr("comment_count+ ?", 1)).Error
	return
}
func (u *ArticleService) DelArticleCommentCount(id uint) (err error) {
	err = global.DB.Model(&repository.Article{}).Where("id", id).
		UpdateColumn("comment_count", gorm.Expr("comment_count- ?", 1)).Error
	return
}
func (u *ArticleService) AddArticleGoodCount(id uint) (err error) {
	err = global.DB.Model(&repository.Article{}).Where("id", id).
		UpdateColumn("good_count", gorm.Expr("good_count+ ?", 1)).Error
	return
}
func (u *ArticleService) DelArticleGoodCount(id uint) (err error) {
	err = global.DB.Model(&repository.Article{}).Where("id", id).
		UpdateColumn("good_count", gorm.Expr("good_count- ?", 1)).Error
	return
}
