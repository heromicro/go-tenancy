package param

type AliyunOSS struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	BucketName      string `json:"bucketName"`
	BucketUrl       string `json:"bucketUrl"`
}

func GetAliyunOSSConfig() (AliyunOSS, error) {
	config := AliyunOSS{}
	alipayConfigs, err := GetConfigByCateKey("aliyun_oss", 0)
	if err != nil {
		return config, err
	}

	for _, alipayConfig := range alipayConfigs {
		if alipayConfig.ConfigKey == "aliyun_endpoint" {
			config.Endpoint = alipayConfig.Value
		}
		if alipayConfig.ConfigKey == "aliyun_access_key_id" {
			config.AccessKeyId = alipayConfig.Value
		}
		if alipayConfig.ConfigKey == "aliyun_access_key_secret" {
			config.AccessKeySecret = alipayConfig.Value
		}
		if alipayConfig.ConfigKey == "aliyun_bucket_name" {
			config.BucketName = alipayConfig.Value
		}
		if alipayConfig.ConfigKey == "aliyun_bucket_url" {
			config.BucketUrl = alipayConfig.Value
		}
	}
	return config, nil
}
