package service

import (
	"github.com/wenknow/gin-vue-blog/server/global"
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"github.com/wenknow/gin-vue-blog/server/model/request"
)

type TagService struct {
}

func (u *TagService) GetTag(cond map[string]interface{}) (Tag repository.Tag, err error) {
	err = buildCond(cond).First(&Tag).Error
	return
}

func (u *TagService) GetTagList(cond map[string]interface{}) (TagList []repository.Tag, err error) {
	err = buildCond(cond).Find(&TagList).Error
	return
}

func (u *TagService) GetTagByPk(id uint) (Tag repository.Tag, err error) {
	err = global.DB.First(&Tag, id).Error
	return
}
func (u *TagService) GetTagByName(name string) (Tag repository.Tag, err error) {
	err = global.DB.Where("name = ?", name).First(&Tag).Error
	return
}

func (u *TagService) PagerTag(cond map[string]interface{}, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNum - 1)
	db := buildCond(cond).Model(&repository.Tag{})
	var TagList []repository.Tag
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&TagList).Error
	return TagList, total, err
}

func (u *TagService) AddTag(Tag repository.Tag) (uid uint, err error) {
	err = global.DB.Create(&Tag).Error
	return Tag.ID, err
}

func (u *TagService) SaveTag(Tag repository.Tag) (err error) {
	err = global.DB.Save(&Tag).Error
	return
}

func (u *TagService) UpdateTag(Tag repository.Tag) (err error) {
	err = global.DB.Updates(&Tag).Error
	return
}

func (u *TagService) DelTag(id uint) (err error) {
	err = global.DB.Delete(&repository.Tag{}, id).Error
	return
}
