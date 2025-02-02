package service

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/utils/param"
	"github.com/snowlyg/go-tenancy/utils/upload"
	"github.com/snowlyg/multi"
)

// GetMediaMap
func GetMediaMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	file, err := FindFile(id)
	if err != nil {
		return form, err
	}

	formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"name","value":"%s","title":"名称","props":{"type":"text","placeholder":"请输入名称"},"validate":[{"message":"请输入名称","required":true,"type":"string","trigger":"change"}]}],"action":"","method":"POST","title":"编辑配置","config":{}}`, file.Name)

	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}

	form.SetAction(fmt.Sprintf("%s/%d", "/media/updateMediaName", id), ctx)
	return form, err
}

// Upload
func Upload(file model.TenancyMedia) (model.TenancyMedia, error) {
	err := g.TENANCY_DB.Create(&file).Error
	return file, err
}

// FindFiles
func FindFiles(ids []uint) ([]model.TenancyMedia, error) {
	var files []model.TenancyMedia
	err := g.TENANCY_DB.Where("id in ?", ids).Find(&files).Error
	return files, err
}

// FindFile
func FindFile(id uint) (model.TenancyMedia, error) {
	var file model.TenancyMedia
	err := g.TENANCY_DB.Where("id = ?", id).First(&file).Error
	return file, err
}

// DeleteFile
func DeleteFile(ids []uint) error {
	files, err := FindFiles(ids)
	if err != nil {
		return fmt.Errorf("find files %w", err)
	}

	delIds := []uint{}
	for _, file := range files {

		oss := upload.NewOss()
		if err = oss.DeleteFile(file.Key); err != nil {
			continue
		}
		delIds = append(delIds, file.ID)
	}

	err = g.TENANCY_DB.Unscoped().Delete(&model.TenancyMedia{}, delIds).Error
	return err
}

func GetFileRecordInfoList(info request.MediaPageInfo, ctx *gin.Context) (interface{}, int64, error) {
	var total int64
	fileLists := []model.TenancyMedia{}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB
	if multi.IsTenancy(ctx) {
		db = db.Where("sys_tenancy_id = ?", multi.GetTenancyId(ctx))
	}
	if info.Name != "" {
		db = db.Where("name like ?", fmt.Sprintf("%s%%", info.Name))
	}

	err := db.Find(&fileLists).Count(&total).Error
	if err != nil {
		return fileLists, total, err
	}
	db = OrderBy(db, info.OrderBy, info.SortBy)
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	if err != nil {
		return fileLists, total, err
	}
	for i := 0; i < len(fileLists); i++ {
		url := fileLists[i].Url
		if !strings.Contains(url, "http://") && !strings.Contains(url, "https://") {
			seitURL, _ := param.GetSeitURL()
			fileLists[i].Url = fmt.Sprintf("%s/%s", seitURL, fileLists[i].Url)
		}
	}
	return fileLists, total, err
}

// UpdateMediaName
func UpdateMediaName(updateMediaName request.UpdateMediaName, id string) error {
	return g.TENANCY_DB.Model(&model.TenancyMedia{}).Where("id = ?", id).Update("name", updateMediaName.Name).Error
}

// UploadFile
func UploadFile(header *multipart.FileHeader, noSave string, tenancyId uint) (model.TenancyMedia, error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(uploadErr)
	}

	var media model.TenancyMedia
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		media.Url = filePath
		media.Name = header.Filename
		media.Tag = s[len(s)-1]
		media.Key = key
		media.SysTenancyId = tenancyId
		return Upload(media)
	}
	return media, nil
}
