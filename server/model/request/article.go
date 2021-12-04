package request

type ArticleListReq struct {
	PageInfo
	Keyword string `json:"keyword" validate:"" label:"关键词"`
	Tag     string `json:"tag" validate:"" label:"标签"`
}
type MyArticleListReq struct {
	PageInfo
	Uid      uint `json:"uid" validate:"required" label:"用户id"`
	TagId    uint `json:"tag_id" validate:"" label:"标签id"`
	CgId     uint `json:"cg_id" validate:"" label:"分类id"`
	PublicIs bool `json:"public_is" validate:"" label:"是否公开"`
}

type ArticleReq struct {
	Id uint `json:"id" validate:"required" label:"文章id"`
}

type ArticleAboutListReq struct {
	Id uint `json:"id" validate:"required" label:"文章id"`
}

type ArticleSaveReq struct {
	Id      uint   `json:"id" validate:"" label:"文章id"`
	CgId    uint   `json:"cg_id" validate:"" label:"分类id"`
	Title   string `json:"title" validate:"" label:"标题"`
	Content string `json:"content" validate:"" label:"内容"`
	ImgUrl  string `json:"img_url" validate:"" label:"封面图片"`
}

type ArticlePublishReq struct {
	Id       uint   `json:"id" validate:"required" label:"文章id"`
	CgName   string `json:"cg_name" validate:"required" label:"分类名称"`
	Title    string `json:"title" validate:"required" label:"标题"`
	Content  string `json:"content" validate:"required" label:"内容"`
	ImgUrl   string `json:"img_url" validate:"" label:"封面图片"`
	TagIdArr []uint `json:"tag_id_arr" validate:"required" label:"所属标签"`
}

type ArticleDelReq struct {
	Id uint `json:"id" validate:"required" label:"文章id"`
}
