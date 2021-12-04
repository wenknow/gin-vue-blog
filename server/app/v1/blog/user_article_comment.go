package blog

import (
	"github.com/gin-gonic/gin"
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"github.com/wenknow/gin-vue-blog/server/model/request"
	"github.com/wenknow/gin-vue-blog/server/model/response"
	"github.com/wenknow/gin-vue-blog/server/service"
	"github.com/wenknow/gin-vue-blog/server/utils/verify"
)

type UserArticleCommentApi struct {
}

func (u *UserArticleCommentApi) AddUserArticleComment(c *gin.Context) {
	var req request.AddUserArticleCommentReq
	var res response.UserArticleCommentListResp
	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}
	myUid := verify.GetTokenUid(c)
	newUserArticleComment := repository.UserArticleComment{
		Uid:       myUid,
		ArticleId: req.ArticleId,
		Content:   req.Content,
		Ip:        req.Ip,
		Address:   req.Address,
	}
	newId, err := userArticleCommentService.AddUserArticleComment(newUserArticleComment)
	if err != nil {
		response.FailWithMsg("评论出错", err, c)
		return
	}
	if req.ArticleId != 0 {
		if err := articleService.AddArticleCommentCount(req.ArticleId); err != nil {
			response.FailWithMsg("更新评论数出错", err, c)
			return
		}
	}

	comment, _ := userArticleCommentService.GetUserArticleCommentDetail(newId)
	res.Comment = comment

	response.OkWithData(res, c)
}

func (u *UserArticleCommentApi) AddUserArticleReply(c *gin.Context) {
	var req request.AddUserArticleReplyReq
	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}
	myUid := verify.GetTokenUid(c)

	//获取被回复的评论详情
	comment, err := userArticleCommentService.GetUserArticleCommentByPk(req.ToCommentId)
	if err != nil {
		response.FailWithMsg("查找评论出错", err, c)
		return
	}
	var toSupCommentId uint
	if comment.ToSupCommentId != 0 {
		toSupCommentId = comment.ToSupCommentId
	} else {
		toSupCommentId = comment.ID
	}
	newUserArticleComment := repository.UserArticleComment{
		Uid:            myUid,
		ArticleId:      req.ArticleId,
		Content:        req.Content,
		ToCommentId:    req.ToCommentId,
		ToSupCommentId: toSupCommentId,
		ToUid:          comment.Uid,
		Ip:             req.Ip,
		Address:        req.Address,
	}
	newId, err := userArticleCommentService.AddUserArticleComment(newUserArticleComment)
	if err != nil {
		response.FailWithMsg("评论出错", err, c)
		return
	}

	if err := articleService.AddArticleCommentCount(req.ArticleId); err != nil {
		response.FailWithMsg("更新评论数出错", err, c)
		return
	}

	if err := userArticleCommentService.AddUserArticleCommentReplyCount(req.ToCommentId); err != nil {
		response.FailWithMsg("更新评论的回复数出错", err, c)
		return
	}
	reply, _ := userArticleCommentService.GetUserArticleCommentDetail(newId)

	response.OkWithData(reply, c)
}

func (u *UserArticleCommentApi) DelUserArticleComment(c *gin.Context) {
	var req request.DelUserArticleCommentReq

	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}

	myUid := verify.GetTokenUid(c)
	comment, err := userArticleCommentService.GetUserArticleCommentDetail(req.Id)
	if err != nil {
		response.FailWithMsg("查找评论出错", err, c)
		return
	} else if comment.Uid != myUid {
		response.FailWithMsg("无法删除别人的评论", err, c)
		return
	}

	if err = userArticleCommentService.DelUserArticleComment(req.Id); err != nil {
		response.FailWithMsg("删除评论出错", err, c)
		return
	}

	if err := articleService.DelArticleCommentCount(comment.ArticleId); err != nil {
		response.FailWithMsg("更新评论数出错", err, c)
		return
	}

	if comment.ToCommentId != 0 {
		if err := userArticleCommentService.DelUserArticleCommentReplyCount(comment.ToCommentId); err != nil {
			response.FailWithMsg("更新评论的回复数出错", err, c)
			return
		}
	}

	response.Ok(c)
}

func (u *UserArticleCommentApi) GetUserArticleCommentList(c *gin.Context) {
	var req request.GetUserArticleCommentListReq
	var res []response.UserArticleCommentListResp

	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}
	cond := service.InitCond()
	cond["article_id"] = req.ArticleId
	cond["to_comment_id"] = 0
	cond["ORDER"] = "id desc"
	pageInfo := service.InitPageInfo(req.PageNum, req.PageSize)
	commentList, allNum, err := userArticleCommentService.PagerUserArticleCommentList(cond, pageInfo)
	if err != nil {
		response.Ok(c)
		return
	}

	//获取评论id
	var commentIdArr []uint
	for _, v := range commentList {
		commentIdArr = append(commentIdArr, v.ID)
	}
	if len(commentIdArr) > 0 {
		cond = service.InitCond()
		cond["article_id"] = req.ArticleId
		cond["to_sup_comment_id"] = commentIdArr
		cond["ORDER"] = "id desc"
		replyList, err := userArticleCommentService.GetUserArticleReplyList(cond)
		if err != nil && err != response.ErrNoRecord {
			response.FailWithMsg("获取评论数出错", err, c)
			return
		}
		if err == response.ErrNoRecord {
		} else if err != nil {
			response.FailWithMsg("获取评论数出错", err, c)
			return
		}
		for _, v := range commentList {
			tmp := response.UserArticleCommentListResp{Comment: v}
			if err != response.ErrNoRecord {
				for _, v2 := range replyList {
					if v2.ToSupCommentId == v.ID {
						tmp.ReplyList = append(tmp.ReplyList, v2)
					}
				}
			}
			res = append(res, tmp)
		}
	}

	response.OkWithData(response.PageResult{
		List:     res,
		AllNum:   allNum,
		PageNum:  pageInfo.PageNum,
		PageSize: pageInfo.PageSize,
	}, c)

}
