package upload

import (
	"mime/multipart"

	"github.com/snowlyg/go-tenancy/utils/param"
)

type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

func NewOss() OSS {
	uploadType, _ := param.GetConfigValueByKey("upload_type")
	switch uploadType {
	case "local":
		return &Local{}
	case "qiniu":
		config, _ := param.GetQiniuConfig()
		return &Qiniu{Config: config}
	case "tencent-cos":
		config, _ := param.GetTencentCOSConfig()
		return &TencentCOS{Config: config}
	case "aliyun-oss":
		config, _ := param.GetAliyunOSSConfig()
		return &AliyunOSS{Config: config}
	default:
		return &Local{}
	}
}
