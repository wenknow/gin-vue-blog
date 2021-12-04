package service

import (
	"github.com/wenknow/gin-vue-blog/server/global"
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"github.com/wenknow/gin-vue-blog/server/model/request"
)

type ArticleCgService struct {
}

func (u *ArticleCgService) GetArticleCg(cond map[string]interface{}) (articleCg repository.ArticleCg, err error) {
	err = buildCond(cond).First(&articleCg).Error
	return
}

func (u *ArticleCgService) GetArticleCgList(cond map[string]interface{}) (articleCgList []repository.ArticleCg, err error) {
	err = buildCond(cond).Find(&articleCgList).Error
	return
}

func (u *ArticleCgService) GetArticleCgByPk(id uint) (articleCg repository.ArticleCg, err error) {
	err = global.DB.First(&articleCg, id).Error
	return
}

func (u *ArticleCgService) PagerArticleCg(cond map[string]interface{}, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNum - 1)
	db := buildCond(cond).Model(&repository.ArticleCg{})
	var articleCgList []repository.ArticleCg
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&articleCgList).Error
	return articleCgList, total, err
}

func (u *ArticleCgService) AddArticleCg(articleCg repository.ArticleCg) (uid uint, err error) {
	err = global.DB.Create(&articleCg).Error
	return articleCg.ID, err
}

func (u *ArticleCgService) SaveArticleCg(articleCg repository.ArticleCg) (err error) {
	err = global.DB.Save(&articleCg).Error
	return
}

func (u *ArticleCgService) UpdateArticleCg(articleCg repository.ArticleCg) (err error) {
	err = global.DB.Updates(&articleCg).Error
	return
}

func (u *ArticleCgService) DelArticleCg(id uint) (err error) {
	err = global.DB.Delete(&repository.ArticleCg{}, id).Error
	return
}

func (u *ArticleCgService) GetArticleCgByNameUid(name string, uid uint) (articleCg repository.ArticleCg, err error) {
	err = global.DB.Where("name", name).Where("uid", uid).First(&articleCg).Error
	return
}

func (u *ArticleCgService) GetArticleCgListByUid(uid uint) (articleCgList []repository.ArticleCg, err error) {
	err = global.DB.Where("uid", uid).Find(&articleCgList).Error
	return
}
