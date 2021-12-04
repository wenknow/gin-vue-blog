package response

import "github.com/wenknow/gin-vue-blog/server/model/repository"

type UserArticleCommentResp struct {
	repository.UserArticleComment
	UserName      string `json:"user_name" `
	UserHeadUrl   string `json:"user_head_url" `
	ToUserName    string `json:"to_user_name" `
	ToUserHeadUrl string `json:"to_user_head_url" `
}

type UserArticleCommentListResp struct {
	Comment   UserArticleCommentResp   `json:"comment" `
	ReplyList []UserArticleCommentResp `json:"reply_list" `
}
