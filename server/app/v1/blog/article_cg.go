package blog

import (
	"github.com/gin-gonic/gin"
	"github.com/wenknow/gin-vue-blog/server/model/config"
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"github.com/wenknow/gin-vue-blog/server/model/request"
	"github.com/wenknow/gin-vue-blog/server/model/response"
	"github.com/wenknow/gin-vue-blog/server/utils/verify"
)

type ArticleCgApi struct {
}

func (articleCgApi *ArticleCgApi) GetArticleCgList(c *gin.Context) {

	var req request.ArticleCgListReq
	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}
	cgList, err := articleCgService.GetArticleCgListByUid(req.Uid)
	if err != nil {
		response.FailWithMsg("获取分类出错", err, c)
	}
	response.OkWithData(cgList, c)

}

func (articleCgApi *ArticleCgApi) AddArticleCg(c *gin.Context) {
	var req request.ArticleCgReq
	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}
	if _, err := articleCgService.GetArticleCgByNameUid(req.Name, req.Uid); err != response.ErrNoRecord {
		response.FailWithMsg("该分类已存在", nil, c)
		return
	}
	myUid := verify.GetTokenUid(c)
	articleCg := repository.ArticleCg{
		Name: req.Name,
		Sort: req.Sort,
		Uid:  myUid,
	}
	if _, err := articleCgService.AddArticleCg(articleCg); err != nil {
		response.FailWithMsg("添加分类出错", err, c)
	} else {
		response.Ok(c)
	}
}

func (articleCgApi *ArticleCgApi) UpdateArticleCg(c *gin.Context) {
	var req request.ArticleCgUpdateReq
	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}

	myUid := verify.GetTokenUid(c)
	old, err := articleCgService.GetArticleCgByPk(req.Id)
	if err != nil {
		response.FailWithMsg("获取分类失败", err, c)
		return
	}
	if old.Uid != myUid {
		response.FailWithMsg("无法更新分类的文章", err, c)
		return
	}
	articleCg := repository.ArticleCg{
		Model: config.Model{ID: req.Id},
		Name:  req.Name,
		Sort:  req.Sort,
	}
	if err := articleCgService.UpdateArticleCg(articleCg); err != nil {
		response.FailWithMsg("更新分类出错", err, c)
	} else {
		response.Ok(c)
	}
}

func (articleCgApi *ArticleCgApi) DelArticleCg(c *gin.Context) {
	var req request.ArticleCgDelReq
	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}

	myUid := verify.GetTokenUid(c)
	old, err := articleCgService.GetArticleCgByPk(req.Id)
	if err != nil {
		response.FailWithMsg("获取分类失败", err, c)
		return
	}
	if old.Uid != myUid {
		response.FailWithMsg("无法更新分类的文章", err, c)
		return
	}
	if err := articleCgService.DelArticleCg(req.Id); err != nil {
		response.FailWithMsg("删除分类出错", err, c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
