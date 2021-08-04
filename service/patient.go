package service

import (
	"fmt"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
)

func FindOrCreatePatient(patient model.Patient) (model.Patient, error) {
	err := g.TENANCY_DB.Model(&model.Patient{}).
		Where(model.Patient{HospitalNO: patient.HospitalNO}).
		// 更新患者的手机号码，年龄，科室，床号，病种 兼容重复入院
		Assign(model.Patient{Phone: patient.Phone, Age: patient.Age, LocName: patient.LocName, BedNum: patient.BedNum, Disease: patient.Disease}).
		FirstOrCreate(&patient).Error
	if err != nil {
		return patient, fmt.Errorf("find or create patient %w", err)
	}
	return patient, nil
}

// GetPatientInfoList
func GetPatientInfoList(info request.PageInfo, tenancyId uint) ([]response.PatientList, int64, error) {
	var patientList []response.PatientList
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.Patient{}).
		Select("patients.*,sys_tenancies.name as hospital_name").
		Joins("left join sys_tenancies on patients.sys_tenancy_id = sys_tenancies.id").
		Where("patients.sys_tenancy_id = ?", tenancyId)
	err := db.Count(&total).Error
	if err != nil {
		return patientList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&patientList).Error

	return patientList, total, err
}

func GetPatientById(patientId, tenancyId uint) (model.Patient, error) {
	var patient model.Patient
	err := g.TENANCY_DB.Model(&model.Patient{}).Where("id =?", patientId).Where("sys_tenancy_id = ?", tenancyId).First(&patient).Error
	if err != nil {
		return patient, err
	}
	return patient, nil
}
