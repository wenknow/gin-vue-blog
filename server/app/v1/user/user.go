package user

import (
	"github.com/gin-gonic/gin"
	"github.com/wenknow/gin-vue-blog/server/model/request"
	"github.com/wenknow/gin-vue-blog/server/model/response"
	"github.com/wenknow/gin-vue-blog/server/utils/verify"
)

type UserApi struct {
}

func (u *UserApi) GetUserInfo(c *gin.Context) {

	info := verify.GetUserInfo(c)
	userInfo, err := userService.GetUserByPk(info.ID)
	if err != nil {
		response.FailWithMsg("获取用户信息失败", err, c)
		return
	}

	response.OkWithData(userInfo, c)
}

func (u *UserApi) GetAuthorInfo(c *gin.Context) {
	var req request.UserInfoReq
	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}
	articleUserInfo, err := articleService.GetArticleGroupByUid(req.Id)
	if err != nil && err != response.ErrNoRecord {
		response.FailWithMsg("获取用户文章信息失败", err, c)
		return
	}

	response.OkWithData(articleUserInfo, c)
}
