package param

type Qiniu struct {
	Zone          string `json:"zone"`          // 存储区域
	Bucket        string `json:"bucket"`        // 空间名称
	ImgPath       string `json:"imgPath"`       // CDN加速域名
	AccessKey     string `json:"accessKey"`     // accessKey
	SecretKey     string `json:"secretKey"`     // secretKey
	UseHTTPS      bool   `json:"useHttps"`      // 是否使用https
	UseCdnDomains bool   `json:"useCdnDomains"` // 上传是否使用CDN上传加速
}

// "aliyun_oss", "qiniuyun", "tengxun"
func GetQiniuConfig() (Qiniu, error) {
	config := Qiniu{}
	configs, err := GetConfigByCateKey("qiniuyun", 0)
	if err != nil {
		return config, err
	}

	for _, conf := range configs {
		if conf.ConfigKey == "qiniu_zone" {
			config.Zone = conf.Value
		}
		if conf.ConfigKey == "qiniu_bucket" {
			config.Bucket = conf.Value
		}
		if conf.ConfigKey == "qiniu_img_path" {
			config.ImgPath = conf.Value
		}
		if conf.ConfigKey == "qiniu_access_key" {
			config.AccessKey = conf.Value
		}
		if conf.ConfigKey == "qiniu_secret_key" {
			config.SecretKey = conf.Value
		}
		if conf.ConfigKey == "qiniu_bucket_url" {
			config.SecretKey = conf.Value
		}
		if conf.ConfigKey == "qiniu_use_https" {
			if conf.Value == "1" {
				config.UseHTTPS = true
			}
		}
		if conf.ConfigKey == "qiniu_use_cdn_domains" {
			if conf.Value == "1" {
				config.UseCdnDomains = true
			}
		}
	}
	return config, nil
}
