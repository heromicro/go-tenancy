package config

type System struct {
	Level       string `mapstructure:"level" json:"level" yaml:"level"` // debug,release,test
	Env         string `mapstructure:"env" json:"env" yaml:"env"`       // dev , pro
	Addr        int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType      string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	CacheType   string `mapstructure:"cache-type" json:"cacheType" yaml:"cache-type"`
	AdminPreix  string `mapstructure:"admin-preix" json:"adminPreix" yaml:"admin-preix"`
	ClientPreix string `mapstructure:"client-preix" json:"clientPreix" yaml:"client-preix"`
}
