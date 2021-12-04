package response

import "github.com/wenknow/gin-vue-blog/server/model/repository"

type ArticleTagInfo struct {
	repository.ArticleTag
	TagName string `json:"tag_name"`
}
