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
	patientList := []response.PatientList{}
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.Patient{}).
		Select("patients.*,sys_tenancies.name as hospital_name").
		Joins("left join sys_tenancies on patients.sys_tenancy_id = sys_tenancies.id")
	db = CheckTenancyId(db, tenancyId, "patients.")
	err := db.Count(&total).Error
	if err != nil {
		return patientList, total, err
	}
	db = OrderBy(db, info.OrderBy, info.SortBy)
	err = db.Limit(limit).Offset(offset).Find(&patientList).Error

	return patientList, total, err
}

func GetPatientById(patientId, tenancyId uint) (model.Patient, error) {
	var patient model.Patient
	db := g.TENANCY_DB.Model(&model.Patient{}).Where("id =?", patientId)
	db = CheckTenancyId(db, tenancyId, "")
	err := db.First(&patient).Error
	if err != nil {
		return patient, err
	}
	return patient, nil
}

func GetPatientSelect(tenancyId uint) ([]response.SelectOption, error) {
	selects := []response.SelectOption{
		{ID: 0, Name: "请选择"},
	}
	var patientSelects []response.SelectOption
	err := g.TENANCY_DB.Model(&model.Patient{}).
		Select("id,name").
		Find(&patientSelects).Error
	selects = append(selects, patientSelects...)
	return selects, err
}
