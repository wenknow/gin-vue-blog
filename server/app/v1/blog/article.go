package blog

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wenknow/gin-vue-blog/server/model/config"
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"github.com/wenknow/gin-vue-blog/server/model/request"
	"github.com/wenknow/gin-vue-blog/server/model/response"
	"github.com/wenknow/gin-vue-blog/server/service"
	"github.com/wenknow/gin-vue-blog/server/utils/typeopt"
	"github.com/wenknow/gin-vue-blog/server/utils/verify"
	"strings"
	"time"
	"unicode/utf8"
)

type ArticleApi struct {
}

func (articleApi *ArticleApi) GetArticleList(c *gin.Context) {
	var req request.ArticleListReq
	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}
	var idArr []uint
	var idArr1 []uint
	var idArr2 []uint
	var err error
	cond := service.InitCond()
	if req.Keyword != "" {
		idArr1, err = articleSearchService.GetArticleIdArrBySearch(req.Keyword)
		if err != nil {
			response.FailWithMsg("搜索文章出错", err, c)
			return
		}
		idArr = idArr1
	}
	if req.Tag != "" {
		tag, err := tagService.GetTagByName(req.Tag)
		if err != nil {
			response.FailWithMsg("搜索标签出错", err, c)
			return
		}
		idArr2, err = articleTagService.GetArticleIdsByTagId(tag.ID)
		if err != nil {
			response.FailWithMsg("搜索文章标签出错", err, c)
			return
		}
		idArr = idArr2
	}
	if req.Keyword != "" && req.Tag != "" {
		idArr = typeopt.IntersectUint(idArr1, idArr2)
	}
	if req.Keyword != "" || req.Tag != "" {
		if len(idArr) == 0 {
			cond["article.id"] = -1
		} else {
			cond["article.id"] = idArr
		}
	}

	cond["article.public_is"] = true
	cond["ORDER"] = "article.id desc"
	pageInfo := service.InitPageInfo(req.PageNum, req.PageSize)
	articleList, allNum, err := articleService.PagerArticle(cond, pageInfo)
	if err != nil {
		response.FailWithMsg("查找文章出错", err, c)
		return
	}
	urls, err := defaultImgService.GetRandomUrls(pageInfo.PageSize)
	if err != nil {
		response.FailWithMsg("查找默认图片出错", err, c)
		return
	}
	for k, v := range articleList {
		if v.ImgUrl == "" {
			articleList[k].ImgUrl = urls[k]
		}
	}
	response.OkWithData(response.PageResult{
		List:     articleList,
		AllNum:   allNum,
		PageNum:  pageInfo.PageNum,
		PageSize: pageInfo.PageSize,
	}, c)

}

func (articleApi *ArticleApi) GetUserArticleList(c *gin.Context) {
	var req request.MyArticleListReq
	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}
	cond := service.InitCond()
	cond["article.uid"] = req.Uid
	if req.CgId != 0 {
		cond["article.cg_id"] = req.CgId
	}
	cond["article.public_is"] = req.PublicIs
	cond["ORDER"] = "article.id desc"
	pageInfo := service.InitPageInfo(req.PageNum, req.PageSize)
	articleList, allNum, err := articleService.PagerArticle(cond, pageInfo)
	if err != nil {
		response.FailWithMsg("查找文章出错", err, c)
		return
	}
	urls, err := defaultImgService.GetRandomUrls(pageInfo.PageSize)
	if err != nil {
		response.FailWithMsg("查找默认图片出错", err, c)
		return
	}
	for k, v := range articleList {
		if v.ImgUrl == "" {
			articleList[k].ImgUrl = urls[k]
		}
	}
	response.OkWithData(response.PageResult{
		List:     articleList,
		AllNum:   allNum,
		PageNum:  pageInfo.PageNum,
		PageSize: pageInfo.PageSize,
	}, c)

}

func (articleApi *ArticleApi) GetArticleDetail(c *gin.Context) {
	var req request.ArticleReq
	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}
	articleDetail, err := articleService.GetArticleInfoByPk(req.Id)
	if err != nil {
		response.FailWithMsg("查找文章出错", err, c)
		return
	}

	myUid := verify.GetTokenUid(c)
	NewUserArticleBrowse := repository.UserArticleBrowse{
		Uid:       myUid,
		ArticleId: articleDetail.ID,
	}
	if _, err := userArticleBrowseService.AddUserArticleBrowse(NewUserArticleBrowse); err != nil {
		response.FailWithMsg("新增文章浏览记录失败", err, c)
		return
	} else {
		if err := articleService.UpdateArticleBrowseCount(articleDetail.ID); err != nil {
			response.FailWithMsg("更新文章浏览记录失败", err, c)
			return
		}
	}
	var userArticleAction response.UserArticleAction
	if myUid != 0 {
		if _, err := userArticleGoodService.GetUserArticleGoodByUidArticleId(myUid, articleDetail.ID); err == response.ErrNoRecord {
			userArticleAction.IsGood = false
		} else if err != nil {
			response.FailWithMsg("查找点赞记录出错", err, c)
			return
		} else {
			userArticleAction.IsGood = true
		}
	}

	response.OkWithData(response.ArticleDetailResp{
		Detail:     articleDetail,
		UserAction: userArticleAction,
	}, c)
}

func (articleApi *ArticleApi) GetRelatedArticleList(c *gin.Context) {
	var req request.ArticleAboutListReq
	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}
	articleTagList, err := articleTagService.GetArticleTagListByArticleId(req.Id)
	if err != nil {
		response.FailWithMsg("查找文章标签出错", err, c)
		return
	}
	var tagArr []uint
	for _, tag := range articleTagList {
		tagArr = append(tagArr, tag.TagId)
	}

	cond := service.InitCond()
	cond["LIMIT"] = 5
	cond["ORDER"] = "browse_count desc, good_count desc"
	list, err := articleService.GetArticleList(cond)
	if err != nil {
		response.FailWithMsg("查找相关文章出错", err, c)
		return
	}
	response.OkWithData(list, c)
}

func (articleApi *ArticleApi) SaveArticle(c *gin.Context) {
	var req request.ArticleSaveReq
	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}
	userInfo := verify.GetUserInfo(c)

	article := repository.Article{
		Title:    req.Title,
		CgId:     req.CgId,
		Content:  req.Content,
		ImgUrl:   req.ImgUrl,
		Uid:      userInfo.ID,
		PublicIs: false,
	}
	if req.Id != 0 {
		article.ID = req.Id
		old, err := articleService.GetArticleByPk(req.Id)
		if err != nil {
			response.FailWithMsg("获取文章失败", err, c)
			return
		}
		if old.Uid != userInfo.ID {
			response.FailWithMsg("无法更新别人的文章", err, c)
			return
		}
		if err := articleService.UpdateArticle(article); err != nil {
			response.FailWithMsg("更新文章草稿出错", err, c)
			return
		}
	} else {
		if article.Title == "" {
			t := time.Now()
			str := fmt.Sprintf("%d%d%d%d%d%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
			article.Title = "未命名-" + str
		}
		if newId, err := articleService.AddArticle(article); err != nil {
			response.FailWithMsg("新增文章草稿出错", err, c)
			return
		} else {
			article.ID = newId
		}
	}
	response.OkWithData(article.ID, c)
}

func (articleApi *ArticleApi) PublishArticle(c *gin.Context) {
	var req request.ArticlePublishReq
	_ = c.ShouldBindJSON(&req)
	fmt.Println(c.Request.Body)
	fmt.Println(req)

	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}
	var cgId uint
	userInfo := verify.GetUserInfo(c)
	//获取分类
	myCg, err := articleCgService.GetArticleCgByNameUid(req.CgName, userInfo.ID)
	if err == response.ErrNoRecord {
		newCg := repository.ArticleCg{
			Name: req.CgName,
			Uid:  userInfo.ID,
		}
		if cgId, err = articleCgService.AddArticleCg(newCg); err != nil {
			response.FailWithMsg("新增分类失败", err, c)
			return
		}
	} else if err != nil {
		response.FailWithMsg("查找分类出错", err, c)
		return
	} else {
		cgId = myCg.ID
	}

	re := []string{" ", "\n", "#", "`"}
	contentStr := typeopt.ReplaceStr(req.Content, re)
	wordCount := utf8.RuneCountInString(contentStr)
	desc := contentStr
	if wordCount > 100 {
		desc = string([]rune(contentStr)[:100]) + "..."
	}

	old, err := articleService.GetArticleByPk(req.Id)
	if err != nil {
		response.FailWithMsg("获取文章失败", err, c)
		return
	}
	if old.Uid != userInfo.ID {
		response.FailWithMsg("无法发布别人的文章", err, c)
		return
	}

	var article = repository.Article{
		Model:     config.Model{ID: req.Id},
		Title:     req.Title,
		CgId:      cgId,
		Content:   req.Content,
		ImgUrl:    req.ImgUrl,
		Uid:       userInfo.ID,
		PublicIs:  true,
		TopIs:     false,
		Desc:      desc,
		WordCount: uint(wordCount),
	}
	if err := articleService.UpdateArticle(article); err != nil {
		response.FailWithMsg("发布出错", err, c)
	}

	//处理标签数据
	for _, tagId := range req.TagIdArr {
		if _, err := tagService.GetTagByPk(tagId); err != nil {
			response.FailWithMsg("获取标签失败", err, c)
			return
		}

		if _, err := articleTagService.GetArticleTagByTagArticle(article.ID, tagId); err == response.ErrNoRecord {
			var newArticleTag = repository.ArticleTag{
				TagId:     tagId,
				ArticleId: article.ID,
				Uid:       userInfo.ID,
			}
			if _, err := articleTagService.AddArticleTag(newArticleTag); err != nil {
				response.FailWithMsg("新增文章标签失败", err, c)
				return
			}
		} else if err != nil {
			response.FailWithMsg("获取文章标签失败", err, c)
			return
		}
	}
	cond := service.InitCond()
	cond["article_id"] = article.ID
	articleTagList, err := articleTagService.GetArticleTagList(cond)
	if err != nil {
		response.FailWithMsg("获取文章标签失败", err, c)
		return
	}
	for _, articleTag := range articleTagList {
		if _, found := typeopt.InUintArray(req.TagIdArr, articleTag.TagId); !found {
			if err := articleTagService.DelArticleTag(articleTag.ID); err != nil {
				response.FailWithMsg("删除文章标签失败", err, c)
				return
			}
		}
	}
	//更新文章标签
	generateTags(article.ID)

	//保存搜索数据
	search := req.Title + contentStr
	var articleSearch = repository.ArticleSearch{
		DelModel: config.DelModel{ID: req.Id},
		Content:  search,
	}

	if _, err := articleSearchService.GetArticleSearchByPk(req.Id); err == response.ErrNoRecord {
		if _, err := articleSearchService.AddArticleSearch(articleSearch); err != nil {
			response.FailWithMsg("新增文章搜索失败", err, c)
			return
		}
	} else if err != nil {
		response.FailWithMsg("获取文章搜索失败", err, c)
		return
	} else {
		if err := articleSearchService.UpdateArticleSearch(articleSearch); err != nil {
			response.FailWithMsg("更新文章搜索失败", err, c)
			return
		}
	}

	response.Ok(c)
}

func (articleApi *ArticleApi) DelArticle(c *gin.Context) {
	var req request.ArticleDelReq
	_ = c.ShouldBindJSON(&req)
	if err := verify.Validate.Struct(req); err != nil {
		response.FailWithMsg(verify.Translate(err), nil, c)
		return
	}
	userInfo := verify.GetUserInfo(c)
	old, err := articleService.GetArticleByPk(req.Id)
	if err != nil {
		response.FailWithMsg("获取文章失败", err, c)
		return
	}
	if old.Uid != userInfo.ID {
		response.FailWithMsg("无法删除别人的文章", err, c)
		return
	}
	if err := articleService.DelArticle(req.Id); err != nil {
		response.FailWithMsg("删除出错", err, c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (articleApi *ArticleApi) GenerateDesc(c *gin.Context) {
	cond := service.InitCond()
	list, _ := articleService.GetArticleList(cond)

	for _, article := range list {

		re := []string{" ", "\n", "#", "`"}
		contentStr := typeopt.ReplaceStr(article.Content, re)
		wordCount := utf8.RuneCountInString(contentStr)

		desc := contentStr
		if wordCount > 100 {
			desc = string([]rune(contentStr)[:100]) + "..."
		}

		var article = repository.Article{
			Model:     config.Model{ID: article.ID},
			Desc:      desc,
			WordCount: uint(wordCount),
		}
		if err := articleService.UpdateArticle(article); err != nil {
			response.FailWithMsg("发布出错", err, c)
		}

	}
}

func (articleApi *ArticleApi) GenerateTags(c *gin.Context) {

	cond := service.InitCond()
	articleList, _ := articleService.GetArticleList(cond)
	for _, article := range articleList {
		generateTags(article.ID)
	}
	response.Ok(c)

}

func (articleApi *ArticleApi) GetInfo(c *gin.Context) {

	response.Ok(c)

}

func generateTags(articleId uint) {
	list, _ := articleTagService.GetArticleTagListByArticleId(articleId)
	var s []string
	for _, v := range list {
		s = append(s, v.TagName)
	}
	n := repository.Article{}
	n.ID = articleId
	n.Tags = strings.Join(s, ",")
	_ = articleService.UpdateArticle(n)
}
