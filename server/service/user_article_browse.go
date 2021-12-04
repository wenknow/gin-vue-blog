package service

import (
	"github.com/wenknow/gin-vue-blog/server/global"
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"github.com/wenknow/gin-vue-blog/server/model/request"
)

type UserArticleBrowseService struct {
}

func (u *UserArticleBrowseService) GetUserArticleBrowse(cond map[string]interface{}) (UserArticleBrowse repository.UserArticleBrowse, err error) {
	err = buildCond(cond).First(&UserArticleBrowse).Error
	return
}

func (u *UserArticleBrowseService) GetUserArticleBrowseList(cond map[string]interface{}) (UserArticleBrowseList []repository.UserArticleBrowse, err error) {
	err = buildCond(cond).Find(&UserArticleBrowseList).Error
	return
}

func (u *UserArticleBrowseService) GetUserArticleBrowseByPk(id uint) (UserArticleBrowse repository.UserArticleBrowse, err error) {
	err = global.DB.First(&UserArticleBrowse, id).Error
	return
}

func (u *UserArticleBrowseService) PagerUserArticleBrowse(cond map[string]interface{}, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNum - 1)
	db := buildCond(cond).Model(&repository.UserArticleBrowse{})
	var UserArticleBrowseList []repository.UserArticleBrowse
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&UserArticleBrowseList).Error
	return UserArticleBrowseList, total, err
}

func (u *UserArticleBrowseService) AddUserArticleBrowse(UserArticleBrowse repository.UserArticleBrowse) (uid uint, err error) {
	err = global.DB.Create(&UserArticleBrowse).Error
	return UserArticleBrowse.ID, err
}

func (u *UserArticleBrowseService) SaveUserArticleBrowse(UserArticleBrowse repository.UserArticleBrowse) (err error) {
	err = global.DB.Save(&UserArticleBrowse).Error
	return
}

func (u *UserArticleBrowseService) UpdateUserArticleBrowse(UserArticleBrowse repository.UserArticleBrowse) (err error) {
	err = global.DB.Updates(&UserArticleBrowse).Error
	return
}

func (u *UserArticleBrowseService) DelUserArticleBrowse(id uint) (err error) {
	err = global.DB.Delete(&repository.UserArticleBrowse{}, id).Error
	return
}
