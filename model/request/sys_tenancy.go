package request

import (
	"github.com/snowlyg/go-tenancy/model"
)

type CreateTenancy struct {
	model.SysTenancy
	Username string `json:"username"  binding:"required"`
}

type SetRegionCode struct {
	Id            float64 `json:"id" form:"id" binding:"required,gt=0"`
	SysRegionCode int     `json:"sysRegionCode" binding:"required"`
}

type TenancyPageInfo struct {
	PageInfo
	Date    string `json:"date" form:"date"`
	Status  string `json:"status" form:"status"`
	Keyword string `json:"keyword" form:"keyword"`
}

type UpdateClientTenancy struct {
	Avatar string `json:"avatar"`
	Banner string `json:"banner"`
	Info   string `json:"info"`
	State  int    `json:"state"`
	Tele   string `json:"tele"`
}

type SetCopyProductNum struct {
	CopyNum int `json:"copyNum"`
	Num     int `json:"num"`
	Type    int `json:"type" binding:"required"` // 1:+ ,2:-
}

type LoginDevice struct {
	UUID       string `json:"uuid" binding:"required"`
	Phone      string `json:"phone" form:"phone" binding:"required"`
	Name       string `json:"name" form:"name" binding:"required"`
	Sex        int    `json:"sex" form:"sex"`
	Age        int    `json:"age" form:"age"`
	LocName    string `json:"locName" form:"locName"  binding:"required"`
	BedNum     string `json:"bedNum" form:"bedNum" binding:"required"`
	HospitalNO string `json:"hospitalNo" form:"hospitalNo" binding:"required"`
	Disease    string `json:"disease" form:"disease"  binding:"required"`
}
