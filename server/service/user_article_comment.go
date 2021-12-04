package service

import (
	"github.com/wenknow/gin-vue-blog/server/global"
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"github.com/wenknow/gin-vue-blog/server/model/request"
	"github.com/wenknow/gin-vue-blog/server/model/response"
	"gorm.io/gorm"
)

type UserArticleCommentService struct {
}

func (u *UserArticleCommentService) GetUserArticleCommentDetail(id uint) (detail response.UserArticleCommentResp, err error) {

	err = global.DB.Table("user_article_comment").
		Joins("left join user B ON user_article_comment.uid = B.id").
		Joins("left join user C ON user_article_comment.to_uid = C.id").
		Select("user_article_comment.*, B.name as user_name, B.head_url as user_head_url, C.name as to_user_name, C.head_url as to_user_head_url").
		Where("user_article_comment.id", id).First(&detail).Error
	return detail, err
}

func (u *UserArticleCommentService) PagerUserArticleCommentList(cond map[string]interface{}, info request.PageInfo) (
	list []response.UserArticleCommentResp, total int64, err error) {

	limit := info.PageSize
	offset := info.PageSize * (info.PageNum - 1)
	db := buildCond(cond).Table("user_article_comment").
		Joins("left join user B ON user_article_comment.uid = B.id").
		Joins("left join user C ON user_article_comment.to_uid = C.id")
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Select("user_article_comment.*, B.name as user_name, B.head_url as user_head_url, C.name as to_user_name, C.head_url as to_user_head_url").
		Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err
}

func (u *UserArticleCommentService) GetUserArticleReplyList(cond map[string]interface{}) (UserArticleCommentList []response.UserArticleCommentResp, err error) {
	err = buildCond(cond).Table("user_article_comment").
		Select("user_article_comment.*, B.name as user_name, B.head_url as user_head_url, C.name as to_user_name, C.head_url as to_user_head_url").
		Joins("left join user B ON user_article_comment.uid = B.id").
		Joins("left join user C ON user_article_comment.to_uid = C.id").
		Find(&UserArticleCommentList).Error
	return
}

func (u *UserArticleCommentService) GetUserArticleComment(cond map[string]interface{}) (UserArticleComment repository.UserArticleComment, err error) {
	err = buildCond(cond).First(&UserArticleComment).Error
	return
}

func (u *UserArticleCommentService) GetUserArticleCommentList(cond map[string]interface{}) (UserArticleCommentList []repository.UserArticleComment, err error) {
	err = buildCond(cond).Find(&UserArticleCommentList).Error
	return
}

func (u *UserArticleCommentService) GetUserArticleCommentByPk(id uint) (UserArticleComment repository.UserArticleComment, err error) {
	err = global.DB.First(&UserArticleComment, id).Error
	return
}

func (u *UserArticleCommentService) PagerUserArticleComment(cond map[string]interface{}, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNum - 1)
	db := buildCond(cond).Model(&repository.UserArticleComment{})
	var UserArticleCommentList []repository.UserArticleComment
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&UserArticleCommentList).Error
	return UserArticleCommentList, total, err
}

func (u *UserArticleCommentService) AddUserArticleComment(UserArticleComment repository.UserArticleComment) (uid uint, err error) {
	err = global.DB.Create(&UserArticleComment).Error
	return UserArticleComment.ID, err
}

func (u *UserArticleCommentService) SaveUserArticleComment(UserArticleComment repository.UserArticleComment) (err error) {
	err = global.DB.Save(&UserArticleComment).Error
	return
}

func (u *UserArticleCommentService) UpdateUserArticleComment(UserArticleComment repository.UserArticleComment) (err error) {
	err = global.DB.Updates(&UserArticleComment).Error
	return
}

func (u *UserArticleCommentService) DelUserArticleComment(id uint) (err error) {
	err = global.DB.Delete(&repository.UserArticleComment{}, id).Error
	return
}

func (u *UserArticleCommentService) AddUserArticleCommentReplyCount(id uint) (err error) {
	err = global.DB.Model(&repository.UserArticleComment{}).Where("id", id).
		UpdateColumn("replay_count", gorm.Expr("replay_count+ ?", 1)).Error
	return
}

func (u *UserArticleCommentService) DelUserArticleCommentReplyCount(id uint) (err error) {
	err = global.DB.Model(&repository.UserArticleComment{}).Where("id", id).
		UpdateColumn("replay_count", gorm.Expr("replay_count- ?", 1)).Error
	return
}
