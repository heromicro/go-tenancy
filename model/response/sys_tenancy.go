package response

import "github.com/snowlyg/go-tenancy/model"

type SysTenancy struct {
	TenancyResponse
	model.BaseTenancy
	Username string `json:"username"`
}

type Counts struct {
	Invalid int
	Valid   int
}

type LoginTenancy struct {
	Admin SysAdminUser `json:"admin"`
	Exp   int64        `json:"exp"`
	Token string       `json:"token"`
	Url   string       `json:"url"`
}

type TenancyInfo struct {
	Avatar string `json:"avatar"`
	Banner string `json:"banner"`
	Id     uint   `json:"id"`
	Info   string `json:"info"`
	Name   string `json:"name"`
}
