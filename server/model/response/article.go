package response

import (
	"github.com/wenknow/gin-vue-blog/server/model/repository"
	"time"
)

type ArticleDetailResp struct {
	Detail     ArticleInfo       `json:"detail"`
	UserAction UserArticleAction `json:"user_action"`
}

type ArticleInfo struct {
	repository.Article
	CgName string `json:"cg_name"`
}

type ArticleListInfo struct {
	Id            uint      `json:"id" `
	BrowseCount   uint      `json:"browse_count" `
	CollectCount  uint      `json:"collect_count"`
	CommentCount  uint      `json:"comment_count" `
	WordCount     uint      `json:"word_count"`
	GoodCount     uint      `json:"good_count"`
	Title         string    `json:"title"`
	Desc          string    `json:"desc"`
	ImgUrl        string    `json:"img_url"`
	TopIs         bool      `json:"top_is"`
	Uid           uint      `json:"uid"`
	CreatedAt     time.Time `json:"created_at" `
	UpdatedAt     time.Time `json:"updated_at" `
	AuthorName    string    `json:"author_name"`
	AuthorHeadUrl string    `json:"author_head_url"`
	Tags          string    `json:"tags"`
	CgName        string    `json:"cg_name"`
}

type ArticleGroupByUserInfo struct {
	Id              uint   `json:"id"`
	Name            string `json:"name"`
	HeadUrl         string `json:"head_url"`
	Desc            string `json:"desc"`
	Gender          int8   `json:"gender"`
	AllArticleCount uint   `json:"all_article_count"`
	AllBrowseCount  uint   `json:"all_browse_count"`
	AllGoodCount    uint   `json:"all_good_count"`
	AllCollectCount uint   `json:"all_collect_count"`
	AllCommentCount uint   `json:"all_comment_count"`
}

type UserArticleAction struct {
	IsGood bool `json:"is_good"`
}
