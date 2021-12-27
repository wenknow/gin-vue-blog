package request

type LoginReq struct {
	Email    string `json:"email" validate:"required,email" label:"邮箱地址"`
	Password string `json:"password" validate:"" label:"密码"`
	Code     string `json:"code" form:"code" validate:"" label:"验证编码"`
}

type RegisterReq struct {
	Name     string `json:"name" validate:"required" label:"昵称"`
	Password string `json:"password" validate:"required" label:"密码"`
	Email    string `json:"email" validate:"required,email" label:"邮箱地址"`
	Desc     string `json:"desc" validate:"" label:"简介"`
}

type RegisterCodeReq struct {
	Code string `json:"code" form:"code" validate:"required" label:"验证编码"`
}

type SendEmailCodeReq struct {
	Email string `json:"email" validate:"required,email" label:"邮箱地址"`
}

type GitHubLoginReq struct {
	Code string `json:"code" validate:"required" label:"编码"`
}
