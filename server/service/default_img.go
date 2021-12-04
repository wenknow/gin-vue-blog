package service

import (
	"github.com/wenknow/gin-vue-blog/server/global"
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"github.com/wenknow/gin-vue-blog/server/model/request"
)

type DefaultImgService struct {
}

func (u *DefaultImgService) GetRandomUrls(count int) (urls []string, err error) {
	var DefaultImgList []repository.DefaultImg
	err = global.DB.Select("url").Order("RAND()").Limit(count).Find(&DefaultImgList).Error
	for _, v := range DefaultImgList {
		urls = append(urls, v.Url)
	}
	return
}

func (u *DefaultImgService) GetDefaultImg(cond map[string]interface{}) (DefaultImg repository.DefaultImg, err error) {
	err = buildCond(cond).First(&DefaultImg).Error
	return
}

func (u *DefaultImgService) GetDefaultImgList(cond map[string]interface{}) (DefaultImgList []repository.DefaultImg, err error) {
	err = buildCond(cond).Find(&DefaultImgList).Error
	return
}

func (u *DefaultImgService) GetDefaultImgByPk(id uint) (DefaultImg repository.DefaultImg, err error) {
	err = global.DB.First(&DefaultImg, id).Error
	return
}

func (u *DefaultImgService) PagerDefaultImg(cond map[string]interface{}, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNum - 1)
	db := buildCond(cond).Model(&repository.DefaultImg{})
	var DefaultImgList []repository.DefaultImg
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&DefaultImgList).Error
	return DefaultImgList, total, err
}

func (u *DefaultImgService) AddDefaultImg(DefaultImg repository.DefaultImg) (uid uint, err error) {
	err = global.DB.Create(&DefaultImg).Error
	return DefaultImg.ID, err
}

func (u *DefaultImgService) SaveDefaultImg(DefaultImg repository.DefaultImg) (err error) {
	err = global.DB.Save(&DefaultImg).Error
	return
}

func (u *DefaultImgService) UpdateDefaultImg(DefaultImg repository.DefaultImg) (err error) {
	err = global.DB.Updates(&DefaultImg).Error
	return
}

func (u *DefaultImgService) DelDefaultImg(id uint) (err error) {
	err = global.DB.Delete(&repository.DefaultImg{}, id).Error
	return
}
