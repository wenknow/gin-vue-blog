package blog

import (
	"github.com/gin-gonic/gin"
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"github.com/wenknow/gin-vue-blog/server/model/request"
	"github.com/wenknow/gin-vue-blog/server/model/response"
	"github.com/wenknow/gin-vue-blog/server/utils/verify"
)

type UserArticleGoodApi struct {
}

func (u *UserArticleGoodApi) AddUserArticleGood(c *gin.Context) {
	var req request.UpdateUserArticleGoodReq
	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}
	myUid := verify.GetTokenUid(c)
	if _, err := userArticleGoodService.GetUserArticleGoodByUidArticleId(myUid, req.ArticleId); err == response.ErrNoRecord {
		newUserArticleGood := repository.UserArticleGood{
			Uid:       myUid,
			ArticleId: req.ArticleId,
		}
		if _, err := userArticleGoodService.AddUserArticleGood(newUserArticleGood); err != nil {
			response.FailWithMsg("点赞出错", err, c)
			return
		}
		if err := articleService.AddArticleGoodCount(req.ArticleId); err != nil {
			response.FailWithMsg("点赞数出错", err, c)
			return
		}

	} else if err != nil {
		response.FailWithMsg("查找点赞记录出错", err, c)
		return
	} else {
		response.FailWithMsg("您已经点过赞了", err, c)
		return
	}

	response.Ok(c)
}

func (u *UserArticleGoodApi) DelUserArticleGood(c *gin.Context) {
	var req request.UpdateUserArticleGoodReq
	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}
	myUid := verify.GetTokenUid(c)
	if userArticleGood, err := userArticleGoodService.GetUserArticleGoodByUidArticleId(myUid, req.ArticleId); err == response.ErrNoRecord {
		response.FailWithMsg("您已经取消过赞了", err, c)
		return
	} else if err != nil {
		response.FailWithMsg("查找点赞记录出错", err, c)
		return
	} else {
		if err := userArticleGoodService.DelUserArticleGood(userArticleGood.ID); err != nil {
			response.FailWithMsg("取消点赞出错", err, c)
			return
		}
		if err := articleService.DelArticleGoodCount(req.ArticleId); err != nil {
			response.FailWithMsg("点赞数出错", err, c)
			return
		}
	}

	response.Ok(c)
}
