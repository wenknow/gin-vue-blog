package service

import (
	"github.com/wenknow/gin-vue-blog/server/global"
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"github.com/wenknow/gin-vue-blog/server/model/request"
)

type ArticleSearchService struct {
}

func (u *ArticleSearchService) GetArticleIdArrBySearch(keyword string) (idArr []uint, err error) {
	var articleSearchList []repository.ArticleSearch
	err = global.DB.Select("id").Where("content LIKE ?", "%"+keyword+"%").Find(&articleSearchList).Error
	if err != nil {
		return idArr, err
	}
	for _, value := range articleSearchList {
		idArr = append(idArr, value.ID)
	}
	return idArr, nil
}

func (u *ArticleSearchService) GetArticleSearch(cond map[string]interface{}) (ArticleSearch repository.ArticleSearch, err error) {
	err = buildCond(cond).First(&ArticleSearch).Error
	return
}

func (u *ArticleSearchService) GetArticleSearchList(cond map[string]interface{}) (ArticleSearchList []repository.ArticleSearch, err error) {
	err = buildCond(cond).Find(&ArticleSearchList).Error
	return
}

func (u *ArticleSearchService) GetArticleSearchByPk(id uint) (ArticleSearch repository.ArticleSearch, err error) {
	err = global.DB.First(&ArticleSearch, id).Error
	return
}

func (u *ArticleSearchService) PagerArticleSearch(cond map[string]interface{}, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNum - 1)
	db := buildCond(cond).Model(&repository.ArticleSearch{})
	var ArticleSearchList []repository.ArticleSearch
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&ArticleSearchList).Error
	return ArticleSearchList, total, err
}

func (u *ArticleSearchService) AddArticleSearch(ArticleSearch repository.ArticleSearch) (uid uint, err error) {
	err = global.DB.Create(&ArticleSearch).Error
	return ArticleSearch.ID, err
}

func (u *ArticleSearchService) SaveArticleSearch(ArticleSearch repository.ArticleSearch) (err error) {
	err = global.DB.Save(&ArticleSearch).Error
	return
}

func (u *ArticleSearchService) UpdateArticleSearch(ArticleSearch repository.ArticleSearch) (err error) {
	err = global.DB.Updates(&ArticleSearch).Error
	return
}

func (u *ArticleSearchService) DelArticleSearch(id uint) (err error) {
	err = global.DB.Delete(&repository.ArticleSearch{}, id).Error
	return
}
