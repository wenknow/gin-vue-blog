package response

import "github.com/wenknow/gin-vue-blog/server/model/repository"

type LoginResponse struct {
	User      repository.User `json:"user"`
	Token     string          `json:"token"`
	ExpiresAt int64           `json:"expires_at"`
}

type UserInfoResponse struct {
	Info    repository.User        `json:"info"`
	Article ArticleGroupByUserInfo `json:"article"`
}
