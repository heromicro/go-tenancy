package upload

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/utils/param"
	"go.uber.org/zap"
)

type AliyunOSS struct {
	Config param.AliyunOSS
}

func (ao *AliyunOSS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	bucket, err := NewBucket(ao)
	if err != nil {
		g.TENANCY_LOG.Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	// 读取本地文件。
	f, openError := file.Open()
	if openError != nil {
		g.TENANCY_LOG.Error("function file.Open() Failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Failed, err:" + openError.Error())
	}

	//上传阿里云路径 文件名格式 自己可以改 建议保证唯一性
	yunFileTmpPath := filepath.Join("uploads", time.Now().Format("2006-01-02")) + "/" + file.Filename

	// 上传文件流。
	err = bucket.PutObject(yunFileTmpPath, f)
	if err != nil {
		g.TENANCY_LOG.Error("function formUploader.PUT() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function formUploader.PUT() Failed, err:" + err.Error())
	}

	return ao.Config.BucketUrl + "/" + yunFileTmpPath, yunFileTmpPath, nil
}

func (ao *AliyunOSS) DeleteFile(key string) error {
	bucket, err := NewBucket(ao)
	if err != nil {
		g.TENANCY_LOG.Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err = bucket.DeleteObject(key)
	if err != nil {
		g.TENANCY_LOG.Error("function bucketManager.DELETE() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.DELETE() Filed, err:" + err.Error())
	}

	return nil
}

func NewBucket(ao *AliyunOSS) (*oss.Bucket, error) {
	// 创建OSSClient实例。
	client, err := oss.New(ao.Config.Endpoint, ao.Config.AccessKeyId, ao.Config.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(ao.Config.BucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
