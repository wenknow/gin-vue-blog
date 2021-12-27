package config

type GithubLogin struct {
	UserUrl      string `mapstructure:"user-url" json:"userUrl" yaml:"user-url"`                // 授权地址
	AuthUrl      string `mapstructure:"auth-url" json:"authUrl" yaml:"auth-url"`                // 授权地址
	ClientId     string `mapstructure:"client-id" json:"clientId" yaml:"client-id"`             // 客户id
	ClientSecret string `mapstructure:"client-secret" json:"clientSecret" yaml:"client-secret"` // 客户密钥
}
