package config

type WechatPay struct {
	WxPkSerialNo string `mapstructure:"wx-pk-serial-no" json:"wxPkSerialNo" yaml:"wx-pk-serial-no"`
	WxPkContent  string `mapstructure:"wx-pk-content" json:"wxPkContent" yaml:"wx-pk-content"`
}
