package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/utils/param"
	"go.uber.org/zap"
)

type Qiniu struct {
	Config param.Qiniu
}

func (qin *Qiniu) UploadFile(file *multipart.FileHeader) (string, string, error) {
	putPolicy := storage.PutPolicy{Scope: qin.Config.Bucket}
	mac := qbox.NewMac(qin.Config.AccessKey, qin.Config.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := qiniuConfig(qin)
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}

	f, openError := file.Open()
	if openError != nil {
		g.TENANCY_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))

		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	putErr := formUploader.Put(context.Background(), &ret, upToken, fileKey, f, file.Size, &putExtra)
	if putErr != nil {
		g.TENANCY_LOG.Error("function formUploader.PUT() Filed", zap.Any("err", putErr.Error()))
		return "", "", errors.New("function formUploader.PUT() Filed, err:" + putErr.Error())
	}
	return qin.Config.ImgPath + "/" + ret.Key, ret.Key, nil
}

func (qin *Qiniu) DeleteFile(key string) error {
	mac := qbox.NewMac(qin.Config.AccessKey, qin.Config.SecretKey)
	cfg := qiniuConfig(qin)
	bucketManager := storage.NewBucketManager(mac, cfg)
	if err := bucketManager.Delete(qin.Config.Bucket, key); err != nil {
		g.TENANCY_LOG.Error("function bucketManager.DELETE() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.DELETE() Filed, err:" + err.Error())
	}
	return nil
}

func qiniuConfig(qin *Qiniu) *storage.Config {
	cfg := storage.Config{
		UseHTTPS:      qin.Config.UseHTTPS,
		UseCdnDomains: qin.Config.UseCdnDomains,
	}
	switch qin.Config.Zone { // 根据配置文件进行初始化空间对应的机房
	case "ZoneHuadong":
		cfg.Zone = &storage.ZoneHuadong
	case "ZoneHuabei":
		cfg.Zone = &storage.ZoneHuabei
	case "ZoneHuanan":
		cfg.Zone = &storage.ZoneHuanan
	case "ZoneBeimei":
		cfg.Zone = &storage.ZoneBeimei
	case "ZoneXinjiapo":
		cfg.Zone = &storage.ZoneXinjiapo
	}
	return &cfg
}
