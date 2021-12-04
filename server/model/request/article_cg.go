package request

type ArticleCgReq struct {
	Name string `json:"name" validate:"required" label:"分类名称"`
	Uid  uint   `json:"uid" validate:"required" label:"所属用户"`
	Sort int    `json:"sort" validate:"" label:"排序值"`
}
type ArticleCgListReq struct {
	Uid uint `json:"uid" validate:"required" label:"所属用户"`
}

type ArticleCgUpdateReq struct {
	Id   uint   `json:"id" validate:"required" label:"id"`
	Name string `json:"name" validate:"required" label:"分类名称"`
	Sort int    `json:"sort" validate:"" label:"排序值"`
}

type ArticleCgDelReq struct {
	Id uint `json:"id" validate:"required" label:"id"`
}
