package response

import "github.com/snowlyg/go-tenancy/model"

type LoginResponse struct {
	User  interface{} `json:"user"`
	Token string      `json:"AccessToken"`
}

type SysAdminUser struct {
	TenancyResponse
	Username      string `json:"userName"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	NickName      string `json:"nickName"`
	HeaderImg     string `json:"headerImg"`
	AuthorityName string `json:"authorityName"`
	AuthorityType int    `json:"authorityType"`
	Status        int    `json:"status"`
	AuthorityId   string `json:"authorityId"`
	TenancyId     uint   `json:"tenancyId"`
	TenancyName   string `json:"tenancyName"`
	DefaultRouter string `json:"defaultRouter"`
}

type SysGeneralUser struct {
	TenancyResponse
	Username string `json:"userName"`
	model.BaseGeneralInfo
	AuthorityName string `json:"authorityName"`
	AuthorityType int    `json:"authorityType"`
	Status        int    `json:"status"`
	AuthorityId   string `json:"authorityId"`
	TenancyId     uint   `json:"tenancyId"`
	TenancyName   string `json:"tenancyName"`
	DefaultRouter string `json:"defaultRouter"`
}
