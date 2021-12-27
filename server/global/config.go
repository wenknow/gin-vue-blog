package global

import "github.com/wenknow/gin-vue-blog/server/model/config"

type Server struct {
	JWT         config.JWT         `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap         config.Zap         `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis       config.Redis       `mapstructure:"redis" json:"redis" yaml:"redis"`
	Captcha     config.Captcha     `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	System      config.System      `mapstructure:"system" json:"system" yaml:"system"`
	Mysql       config.Mysql       `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Local       config.Local       `mapstructure:"local" json:"local" yaml:"local"`
	Timer       config.Timer       `mapstructure:"timer" json:"timer" yaml:"timer"`
	Email       config.Email       `mapstructure:"email" json:"email" yaml:"email"`
	Encrypt     config.Encrypt     `mapstructure:"encrypt" json:"encrypt" yaml:"encrypt"`
	AliyunOSS   config.AliyunOSS   `mapstructure:"aliyun-oss" json:"aliyunOSS" yaml:"aliyun-oss"`
	GithubLogin config.GithubLogin `mapstructure:"github-login" json:"githubLogin" yaml:"github-login"`
}
