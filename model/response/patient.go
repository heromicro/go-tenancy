package response

import "github.com/snowlyg/go-tenancy/model"

type PatientList struct {
	model.Patient
	HospitalName int `json:"hospitalName"`
}
