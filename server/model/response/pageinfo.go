package response

type PageResult struct {
	List     interface{} `json:"list"`
	AllNum   int64       `json:"all_num"`
	PageNum  int         `json:"page_num"`
	PageSize int         `json:"page_size"`
}
