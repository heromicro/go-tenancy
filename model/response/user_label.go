package response

import "github.com/snowlyg/go-tenancy/model"

type UserLabelWithUserId struct {
	model.UserLabel
	SysUserId uint `json:"sysUserId"`
}
