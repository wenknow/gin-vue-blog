package request

type UserInfoReq struct {
	Id uint `json:"id" validate:"required" label:"用户id"`
}
