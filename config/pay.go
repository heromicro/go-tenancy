package config

type Alipay struct {
	AppId          string `mapstructure:"app-id" json:"appId" yaml:"app-id"`
	PrivateKey     string `mapstructure:"private-key" json:"privateKey" yaml:"private-key"`
	Charset        string `mapstructure:"charset" json:"charset" yaml:"charset"`
	IsProd         bool   `mapstructure:"is-prod" json:"isProd" yaml:"is-prod"`
	SignType       string `mapstructure:"sign-type" json:"signType" yaml:"sign-type"`
	PrivateKeyType string `mapstructure:"private-key-type" json:"privateKeyType" yaml:"private-key-type"`
	NotifyUrl      string `mapstructure:"notify-url" json:"notifyUrl" yaml:"notify-url"`
}
