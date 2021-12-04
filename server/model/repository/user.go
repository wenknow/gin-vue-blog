package repository

import (
	"github.com/wenknow/gin-vue-blog/server/model/config"
)

// User 结构体
type User struct {
	config.Model
	ActivateIs bool   `json:"activate_is" form:"activate_is" gorm:"column:activate_is;comment:是否激活;type:tinyint"`
	Email      string `json:"email" form:"email" gorm:"column:email;comment:邮箱;type:varchar(80);"`
	Gender     int8   `json:"gender" form:"gender" gorm:"column:gender;comment:性别;type:tinyint"`
	Name       string `json:"name" form:"name" gorm:"column:name;comment:名称;type:varchar(40);"`
	Password   string `json:"-" form:"password" gorm:"column:password;comment:密码;type:varchar(80);"`
	Phone      string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;type:varchar(40);"`
	Username   string `json:"username" form:"username" gorm:"column:username;comment:账号;type:varchar(40);"`
	HeadUrl    string `json:"head_url" form:"head_url" gorm:"column:head_url;comment:头像;type:varchar(255);"`
	Desc       string `json:"desc" form:"desc" gorm:"column:desc;comment:简介;type:varchar(255);"`
}
