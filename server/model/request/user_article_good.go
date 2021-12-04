package request

type UpdateUserArticleGoodReq struct {
	ArticleId uint `json:"article_id" validate:"required" label:"文章id"`
}
