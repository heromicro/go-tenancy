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
		if alipayConfig.ConfigKey == "endpoint" {
			config.Endpoint = alipayConfig.Value
		}
		if alipayConfig.ConfigKey == "access_key_id" {
			config.AccessKeyId = alipayConfig.Value
		}
		if alipayConfig.ConfigKey == "access_key_secret" {
			config.AccessKeySecret = alipayConfig.Value
		}
		if alipayConfig.ConfigKey == "bucket_name" {
			config.BucketName = alipayConfig.Value
		}
		if alipayConfig.ConfigKey == "bucket_url" {
			config.BucketUrl = alipayConfig.Value
		}
	}
	return config, nil
}
