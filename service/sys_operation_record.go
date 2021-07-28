package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

// CreateSysOperationRecord 创建记录
func CreateSysOperationRecord(sysOperationRecord model.SysOperationRecord) error {
	return g.TENANCY_DB.Create(&sysOperationRecord).Error
}

// DeleteSysOperationRecordByIds 批量删除记录
func DeleteSysOperationRecordByIds(ids request.IdsReq) error {
	return g.TENANCY_DB.Delete(&[]model.SysOperationRecord{}, "id in (?)", ids.Ids).Error
}

// DeleteSysOperationRecord 删除操作记录
func DeleteSysOperationRecord(sysOperationRecord model.SysOperationRecord) error {
	return g.TENANCY_DB.Delete(&sysOperationRecord).Error
}

// GetSysOperationRecord 根据id获取单条操作记录
func GetSysOperationRecord(id uint) (model.SysOperationRecord, error) {
	var sysOperationRecord model.SysOperationRecord
	err := g.TENANCY_DB.Where("id = ?", id).First(&sysOperationRecord).Error
	return sysOperationRecord, err
}

// GetSysOperationRecordInfoList 分页获取操作记录列表
func GetSysOperationRecordInfoList(info request.SysOperationRecordSearch, ctx *gin.Context) ([]response.SysOperationRecord, int64, error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	var sysOperationRecords []response.SysOperationRecord
	var adminUsers []response.SysAdminUser
	var tenancyUsers []response.SysTenancyUser
	var err error
	db := g.TENANCY_DB.Model(&model.SysOperationRecord{})
	if multi.IsTenancy(ctx) {
		db = db.Where("sys_tenancy_id = ?", multi.GetTenancyId(ctx))
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Method != "" {
		db = db.Where("method LIKE ?", info.Method+"%")
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", info.Path+"%")
	}
	if info.Status > 0 {
		db = db.Where("status = ?", info.Status)
	}
	var total int64
	err = db.Count(&total).Error
	if err != nil {
		return nil, total, err
	}
	err = db.Order("id desc").Limit(limit).Offset(offset).Find(&sysOperationRecords).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, total, err
	}

	var opUserIds []uint
	for _, sysOperationRecord := range sysOperationRecords {
		opUserIds = append(opUserIds, sysOperationRecord.UserID)
	}

	tenancyUsers, err = GetTenancyByUserIds(opUserIds)
	if err != nil {
		return nil, 0, err
	}

	if multi.IsAdmin(ctx) {
		adminUsers, err = GetAdminByUserIds(opUserIds)
		if err != nil {
			return nil, 0, err
		}
	}
	for i := 0; i < len(sysOperationRecords); i++ {
		if len(tenancyUsers) > 0 {
			for _, tenancyUser := range tenancyUsers {
				if tenancyUser.ID == sysOperationRecords[i].UserID {
					sysOperationRecords[i].NickName = tenancyUser.NickName
					sysOperationRecords[i].TenancyName = tenancyUser.TenancyName
					sysOperationRecords[i].UserName = tenancyUser.Username
				}
			}
		}

		if len(adminUsers) > 0 {
			for _, adminUser := range adminUsers {
				if adminUser.ID == sysOperationRecords[i].UserID {
					sysOperationRecords[i].NickName = adminUser.NickName
					sysOperationRecords[i].UserName = adminUser.Username
				}
			}
		}
	}

	return sysOperationRecords, total, err
}
