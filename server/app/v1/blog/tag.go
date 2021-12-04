package blog

import (
	"github.com/gin-gonic/gin"
	"github.com/wenknow/gin-vue-blog/server/model/request"
	"github.com/wenknow/gin-vue-blog/server/model/response"
	"github.com/wenknow/gin-vue-blog/server/service"
	"github.com/wenknow/gin-vue-blog/server/utils/typeopt"
	"github.com/wenknow/gin-vue-blog/server/utils/verify"
	"strings"
)

type TagApi struct {
}

func (tagApi *TagApi) GetMainTagList(c *gin.Context) {
	var req request.MainTagListReq
	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}
	cond := service.InitCond()
	if req.Keyword != "" {
		idArr, err := articleSearchService.GetArticleIdArrBySearch(req.Keyword)
		if err != nil {
			response.FailWithMsg("搜索文章出错", err, c)
			return
		}
		if len(idArr) == 0 {
			cond["A.article_id"] = -1
		} else {
			cond["A.article_id"] = idArr
		}
	}

	data, err := articleTagService.GetArticleTagListGroupByTag(cond)
	if err != nil {
		response.FailWithMsg("获取分类出错", err, c)
		return
	}
	response.OkWithData(data, c)

}

func (tagApi *TagApi) GetUserTagList(c *gin.Context) {

	userInfo := verify.GetUserInfo(c)
	cond := service.InitCond()
	cond["uid"] = userInfo.ID
	tagList, err := articleTagService.GetArticleTagListGroupByTag(cond)
	if err != nil {
		response.FailWithMsg("获取分类出错", err, c)
		return
	}
	response.OkWithData(tagList, c)

}

func (tagApi *TagApi) GetTagList(c *gin.Context) {

	cond := service.InitCond()
	tagList, err := tagService.GetTagList(cond)
	if err != nil {
		response.FailWithMsg("获取分类出错", err, c)
		return
	}
	response.OkWithData(tagList, c)

}

func (tagApi *TagApi) GuessTag(c *gin.Context) {

	var req request.TagGuessReq
	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}
	var ret []uint
	contentStr := req.Content + req.Title
	if contentStr == "" {
		response.OkWithData(ret, c)
		return
	}
	re := []string{" ", "\n", "#", "`"}
	contentStr = typeopt.ReplaceStr(contentStr, re)
	contentStr = strings.ToLower(contentStr)
	cond := service.InitCond()
	tagList, err := tagService.GetTagList(cond)
	if err != nil {
		response.FailWithMsg("获取分类出错", err, c)
	}

	//取前3个出现次数最多的tag
	var sortArr = []int{0, 0, 0}
	var keyArr = []uint{0, 0, 0}
	for _, tag := range tagList {
		count := strings.Count(contentStr, strings.ToLower(tag.Name))
		if count > 0 {
			for i := 0; i < 3; i++ {
				if count > sortArr[i] {
					sortArr = append(sortArr[:i], append(append([]int{count}, sortArr[i:2]...))...)
					keyArr = append(keyArr[:i], append(append([]uint{tag.ID}, keyArr[i:2]...))...)
					break
				}
			}
		}
	}
	for _, v := range keyArr {
		if v != 0 {
			ret = append(ret, v)
		}
	}
	response.OkWithData(ret, c)
}
