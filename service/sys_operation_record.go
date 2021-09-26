package service

import (
	"errors"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"gorm.io/gorm"
)

// CreateSysOperationRecord 创建记录
func CreateSysOperationRecord(sysOperationRecord model.SysOperationRecord) error {
	return g.TENANCY_DB.Create(&sysOperationRecord).Error
}

// GetSysOperationRecordInfoList 分页获取操作记录列表
func GetSysOperationRecordInfoList(info request.SysOperationRecordSearch, tenancyId uint) ([]response.SysOperationRecord, int64, error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	sysOperationRecords := []response.SysOperationRecord{}
	db := g.TENANCY_DB.Model(&model.SysOperationRecord{}).Where("sys_tenancy_id = ?", tenancyId)

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
	err := db.Count(&total).Error
	if err != nil {
		return nil, total, err
	}
	db = OrderBy(db, info.OrderBy, info.SortBy)
	err = db.Order("id desc").Limit(limit).Offset(offset).Find(&sysOperationRecords).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, total, err
	}

	opUserIds := []uint{}
	for _, sysOperationRecord := range sysOperationRecords {
		opUserIds = append(opUserIds, sysOperationRecord.UserID)
	}

	if tenancyId == 0 {
		adminUsers, err := GetAdminByUserIds(opUserIds)
		if err != nil {
			return nil, 0, err
		}

		for i := 0; i < len(sysOperationRecords); i++ {
			if len(adminUsers) > 0 {
				for _, adminUser := range adminUsers {
					if adminUser.ID == sysOperationRecords[i].UserID {
						sysOperationRecords[i].NickName = adminUser.NickName
						sysOperationRecords[i].UserName = adminUser.Username
					}
				}
			}
		}

	} else {
		tenancyUsers, err := GetTenancyByUserIds(opUserIds, tenancyId)
		if err != nil {
			return nil, 0, err
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
		}
	}

	return sysOperationRecords, total, err
}
