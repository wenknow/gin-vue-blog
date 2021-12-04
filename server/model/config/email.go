package config

type Email struct {
	Port   int    `mapstructure:"port" json:"port" yaml:"port"`       // 端口
	Sender string `mapstructure:"sender" json:"sender" yaml:"sender"` // 发送者
	Host   string `mapstructure:"host" json:"host" yaml:"host"`       // 服务器地址
	IsSSL  bool   `mapstructure:"is-ssl" json:"isSSL" yaml:"is-ssl"`  // 是否SSL
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"` // 密钥
	Name   string `mapstructure:"name" json:"name" yaml:"name"`       // 发送者名称
}
