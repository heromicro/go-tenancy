package service

import (
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
)

// CreatePatient
func CreatePatient(patient model.Patient, tenancyId uint) (model.Patient, error) {
	patient.SysTenancyID = tenancyId
	err := g.TENANCY_DB.Create(&patient).Error
	return patient, err
}

// GetPatientByID
func GetPatientByID(id uint) (model.Patient, error) {
	var patient model.Patient
	err := g.TENANCY_DB.Model(&model.Patient{}).
		Where("id = ?", id).
		First(&patient).Error
	return patient, err
}

// DeletePatient
func DeletePatient(id uint) error {
	var product model.Patient
	return g.TENANCY_DB.Where("id = ?", id).Delete(&product).Error
}

// GetPatientInfoList
func GetPatientInfoList(info request.PageInfo) ([]response.PatientList, int64, error) {
	var patientList []response.PatientList
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.Patient{}).
		Select("patients.*,sys_tenancies.name as hospital_name").
		Joins("left join sys_tenancies on patients.sys_tenancy_id = sys_tenancies.id")
	err := db.Count(&total).Error
	if err != nil {
		return patientList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&patientList).Error

	return patientList, total, err
}
