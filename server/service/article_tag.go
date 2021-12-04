package service

import (
	"github.com/wenknow/gin-vue-blog/server/global"
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"github.com/wenknow/gin-vue-blog/server/model/request"
	"github.com/wenknow/gin-vue-blog/server/model/response"
)

type ArticleTagService struct {
}

func (u *ArticleTagService) GetArticleTagListGroupByTag(cond map[string]interface{}) (mainArticleTag []response.MainArticleTag, err error) {
	err = buildCond(cond).Table("article_tag A").Select("count(A.id) as count, A.tag_id, B.name").
		Joins("left join tag B ON A.tag_id = B.id").Group("tag_id").Find(&mainArticleTag).Error
	return
}

func (u *ArticleTagService) GetArticleTag(cond map[string]interface{}) (ArticleTag repository.ArticleTag, err error) {
	err = buildCond(cond).First(&ArticleTag).Error
	return
}

func (u *ArticleTagService) GetArticleTagList(cond map[string]interface{}) (ArticleTagList []repository.ArticleTag, err error) {
	err = buildCond(cond).Find(&ArticleTagList).Error
	return
}

func (u *ArticleTagService) GetArticleTagByPk(id uint) (ArticleTag repository.ArticleTag, err error) {
	err = global.DB.First(&ArticleTag, id).Error
	return
}

func (u *ArticleTagService) PagerArticleTag(cond map[string]interface{}, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNum - 1)
	db := buildCond(cond).Model(&repository.ArticleTag{})
	var ArticleTagList []repository.ArticleTag
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&ArticleTagList).Error
	return ArticleTagList, total, err
}

func (u *ArticleTagService) AddArticleTag(ArticleTag repository.ArticleTag) (uid uint, err error) {
	err = global.DB.Create(&ArticleTag).Error
	return ArticleTag.ID, err
}

func (u *ArticleTagService) SaveArticleTag(ArticleTag repository.ArticleTag) (err error) {
	err = global.DB.Save(&ArticleTag).Error
	return
}

func (u *ArticleTagService) UpdateArticleTag(ArticleTag repository.ArticleTag) (err error) {
	err = global.DB.Updates(&ArticleTag).Error
	return
}

func (u *ArticleTagService) DelArticleTag(id uint) (err error) {
	err = global.DB.Delete(&repository.ArticleTag{}, id).Error
	return
}

func (u *ArticleTagService) GetArticleTagListByArticleId(articleId uint) (articleTagInfoList []response.ArticleTagInfo, err error) {
	err = global.DB.Table("article_tag A").Where("article_id", articleId).Select("A.*, B.name as tag_name").
		Joins("left join tag B ON A.tag_id = B.id").Find(&articleTagInfoList).Error
	return
}

func (u *ArticleTagService) GetArticleTagByTagArticle(articleId uint, tagId uint) (articleTag repository.ArticleTag, err error) {
	err = global.DB.Where("tag_id", tagId).Where("article_id", articleId).First(&articleTag).Error
	return
}

func (u *ArticleTagService) GetArticleIdsByTagId(tagId uint) (articleIdArr []uint, err error) {
	var articleTagList []repository.ArticleTag
	err = global.DB.Where("tag_id", tagId).Find(&articleTagList).Error
	if err != nil {
		return articleIdArr, err
	}
	for _, value := range articleTagList {
		articleIdArr = append(articleIdArr, value.ArticleId)
	}
	return articleIdArr, nil
}
