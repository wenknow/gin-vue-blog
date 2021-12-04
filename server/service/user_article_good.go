package service

import (
	"github.com/wenknow/gin-vue-blog/server/global"
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"github.com/wenknow/gin-vue-blog/server/model/request"
)

type UserArticleGoodService struct {
}

func (u *UserArticleGoodService) GetUserArticleGoodByUidArticleId(uid uint, articleId uint) (UserArticleGood repository.UserArticleGood, err error) {
	err = global.DB.Where("uid", uid).Where("article_id", articleId).First(&UserArticleGood).Error
	return
}

func (u *UserArticleGoodService) GetUserArticleGood(cond map[string]interface{}) (UserArticleGood repository.UserArticleGood, err error) {
	err = buildCond(cond).First(&UserArticleGood).Error
	return
}

func (u *UserArticleGoodService) GetUserArticleGoodList(cond map[string]interface{}) (UserArticleGoodList []repository.UserArticleGood, err error) {
	err = buildCond(cond).Find(&UserArticleGoodList).Error
	return
}

func (u *UserArticleGoodService) GetUserArticleGoodByPk(id uint) (UserArticleGood repository.UserArticleGood, err error) {
	err = global.DB.First(&UserArticleGood, id).Error
	return
}

func (u *UserArticleGoodService) PagerUserArticleGood(cond map[string]interface{}, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNum - 1)
	db := buildCond(cond).Model(&repository.UserArticleGood{})
	var UserArticleGoodList []repository.UserArticleGood
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&UserArticleGoodList).Error
	return UserArticleGoodList, total, err
}

func (u *UserArticleGoodService) AddUserArticleGood(UserArticleGood repository.UserArticleGood) (uid uint, err error) {
	err = global.DB.Create(&UserArticleGood).Error
	return UserArticleGood.ID, err
}

func (u *UserArticleGoodService) SaveUserArticleGood(UserArticleGood repository.UserArticleGood) (err error) {
	err = global.DB.Save(&UserArticleGood).Error
	return
}

func (u *UserArticleGoodService) UpdateUserArticleGood(UserArticleGood repository.UserArticleGood) (err error) {
	err = global.DB.Updates(&UserArticleGood).Error
	return
}

func (u *UserArticleGoodService) DelUserArticleGood(id uint) (err error) {
	err = global.DB.Delete(&repository.UserArticleGood{}, id).Error
	return
}
