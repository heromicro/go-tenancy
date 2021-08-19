package param

type TencentCOS struct {
	Bucket     string `json:"bucket"`
	Region     string `json:"region"`
	SecretID   string `json:"secretID"`
	SecretKey  string `json:"secretKey"`
	BaseURL    string `json:"baseURL"`
	PathPrefix string `json:"pathPrefix"`
}

func GetTencentCOSConfig() (TencentCOS, error) {
	config := TencentCOS{}
	configs, err := GetConfigByCateKey("tengxun", 0)
	if err != nil {
		return config, err
	}

	for _, conf := range configs {
		if conf.ConfigKey == "tengxun_bucket" {
			config.Bucket = conf.Value
		}
		if conf.ConfigKey == "tengxun_region" {
			config.Region = conf.Value
		}
		if conf.ConfigKey == "tengxun_access_key" {
			config.SecretID = conf.Value
		}
		if conf.ConfigKey == "tengxun_secret_key" {
			config.SecretKey = conf.Value
		}
		if conf.ConfigKey == "tengxun_base_url" {
			config.BaseURL = conf.Value
		}
		if conf.ConfigKey == "tengxun_path_prefix" {
			config.PathPrefix = conf.Value
		}
	}
	return config, nil
}
