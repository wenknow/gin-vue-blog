package response

type MainArticleTag struct {
	TagId uint   `json:"tag_id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}
