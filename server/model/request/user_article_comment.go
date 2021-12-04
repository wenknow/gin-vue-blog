package request

type GetUserArticleCommentListReq struct {
	PageInfo
	ArticleId uint `json:"article_id" validate:"" label:"文章id"`
}

type AddUserArticleCommentReq struct {
	ArticleId uint   `json:"article_id" validate:"" label:"文章id"`
	Content   string `json:"content" validate:"required" label:"评论内容"`
	Ip        string `json:"ip" validate:"required" label:"ip"`
	Address   string `json:"address" validate:"required" label:"地址"`
}

type DelUserArticleCommentReq struct {
	Id uint `json:"id" validate:"required" label:"评论id"`
}

type AddUserArticleReplyReq struct {
	ArticleId   uint   `json:"article_id" validate:"required" label:"文章id"`
	ToCommentId uint   `json:"to_comment_id" validate:"required" label:"被回复的评论id"`
	Content     string `json:"content" validate:"required" label:"评论内容"`
	Ip          string `json:"ip" validate:"required" label:"ip"`
	Address     string `json:"address" validate:"required" label:"地址"`
}
