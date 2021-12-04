package request

type PageInfo struct {
	PageNum  int `json:"page_num" form:"page_num"`  // 页码
	PageSize int `json:"page_size" form:"pag_size"` // 每页大小
}
