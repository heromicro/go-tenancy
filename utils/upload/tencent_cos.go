package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/utils/param"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
)

type TencentCOS struct {
	Config param.TencentCOS
}

// UploadFile upload file to COS
func (tcos *TencentCOS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	client := NewClient(tcos)
	f, openError := file.Open()
	if openError != nil {
		g.TENANCY_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)

	_, err := client.Object.Put(context.Background(), tcos.Config.PathPrefix+"/"+fileKey, f, nil)
	if err != nil {
		panic(err)
	}
	return tcos.Config.BaseURL + "/" + tcos.Config.PathPrefix + "/" + fileKey, fileKey, nil
}

// DeleteFile delete file form COS
func (tcos *TencentCOS) DeleteFile(key string) error {
	client := NewClient(tcos)
	name := tcos.Config.PathPrefix + "/" + key
	_, err := client.Object.Delete(context.Background(), name)
	if err != nil {
		g.TENANCY_LOG.Error("function bucketManager.DELETE() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.DELETE() Filed, err:" + err.Error())
	}
	return nil
}

// NewClient init COS client
func NewClient(tcos *TencentCOS) *cos.Client {
	urlStr, _ := url.Parse("https://" + tcos.Config.Bucket + ".cos." + tcos.Config.Region + ".myqcloud.com")
	baseURL := &cos.BaseURL{BucketURL: urlStr}
	client := cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  tcos.Config.SecretID,
			SecretKey: tcos.Config.SecretKey,
		},
	})
	return client
}
