package service

import (
	"github.com/wenknow/gin-vue-blog/server/global"
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"github.com/wenknow/gin-vue-blog/server/model/request"
)

type UserArticleCommentGoodService struct {
}

func (u *UserArticleCommentGoodService) GetUserArticleCommentGood(cond map[string]interface{}) (UserArticleCommentGood repository.UserArticleCommentGood, err error) {
	err = buildCond(cond).First(&UserArticleCommentGood).Error
	return
}

func (u *UserArticleCommentGoodService) GetUserArticleCommentGoodList(cond map[string]interface{}) (UserArticleCommentGoodList []repository.UserArticleCommentGood, err error) {
	err = buildCond(cond).Find(&UserArticleCommentGoodList).Error
	return
}

func (u *UserArticleCommentGoodService) GetUserArticleCommentGoodByPk(id uint) (UserArticleCommentGood repository.UserArticleCommentGood, err error) {
	err = global.DB.First(&UserArticleCommentGood, id).Error
	return
}

func (u *UserArticleCommentGoodService) PagerUserArticleCommentGood(cond map[string]interface{}, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNum - 1)
	db := buildCond(cond).Model(&repository.UserArticleCommentGood{})
	var UserArticleCommentGoodList []repository.UserArticleCommentGood
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&UserArticleCommentGoodList).Error
	return UserArticleCommentGoodList, total, err
}

func (u *UserArticleCommentGoodService) AddUserArticleCommentGood(UserArticleCommentGood repository.UserArticleCommentGood) (uid uint, err error) {
	err = global.DB.Create(&UserArticleCommentGood).Error
	return UserArticleCommentGood.ID, err
}

func (u *UserArticleCommentGoodService) SaveUserArticleCommentGood(UserArticleCommentGood repository.UserArticleCommentGood) (err error) {
	err = global.DB.Save(&UserArticleCommentGood).Error
	return
}

func (u *UserArticleCommentGoodService) UpdateUserArticleCommentGood(UserArticleCommentGood repository.UserArticleCommentGood) (err error) {
	err = global.DB.Updates(&UserArticleCommentGood).Error
	return
}

func (u *UserArticleCommentGoodService) DelUserArticleCommentGood(id uint) (err error) {
	err = global.DB.Delete(&repository.UserArticleCommentGood{}, id).Error
	return
}
