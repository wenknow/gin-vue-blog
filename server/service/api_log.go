package service

import (
	"github.com/wenknow/gin-vue-blog/server/global"
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"github.com/wenknow/gin-vue-blog/server/model/request"
)

type ApiLogService struct {
}

func (u *ApiLogService) GetApiLog(cond map[string]interface{}) (apiLog repository.ApiLog, err error) {
	err = buildCond(cond).First(&apiLog).Error
	return
}

func (u *ApiLogService) GetApiLogList(cond map[string]interface{}) (apiLogList []repository.ApiLog, err error) {
	err = buildCond(cond).Find(&apiLogList).Error
	return
}

func (u *ApiLogService) GetApiLogByPk(id uint) (apiLog repository.ApiLog, err error) {
	err = global.DB.First(&apiLog, id).Error
	return
}

func (u *ApiLogService) PagerApiLog(cond map[string]interface{}, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNum - 1)
	db := buildCond(cond).Model(&repository.ApiLog{})
	var apiLogList []repository.ApiLog
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&apiLogList).Error
	return apiLogList, total, err
}

func (u *ApiLogService) AddApiLog(apiLog repository.ApiLog) (uid uint, err error) {
	err = global.DB.Create(&apiLog).Error
	return apiLog.ID, err
}

func (u *ApiLogService) UpdateApiLog(apiLog repository.ApiLog) (err error) {
	err = global.DB.Save(&apiLog).Error
	return
}

func (u *ApiLogService) DelApiLog(apiLog repository.ApiLog) (err error) {
	err = global.DB.Delete(&apiLog).Error
	return
}
