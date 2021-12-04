package request

type TagGuessReq struct {
	Content string `json:"content" validate:"" label:"文章内容"`
	Title   string `json:"title" validate:"" label:"标题"`
}

type MainTagListReq struct {
	Keyword string `json:"keyword" validate:"" label:"关键词"`
}
