package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"pageSize" form:"pageSize" binding:"required"`
	OrderBy  string `json:"orderBy" form:"orderBy"`
	SortBy   string `json:"sortBy" form:"sortBy"`
}

type ChangeStatus struct {
	Id     uint `json:"id" form:"id" binding:"required,gt=0"`
	Status int  `json:"status" binding:"required"`
}

// Find by id structure
type GetByTenancyId struct {
	TenancyId uint `json:"tenancy_id" uri:"tenancy_id" form:"tenancy_id" binding:"required"`
}

// Find by id structure
type GetById struct {
	Id        uint `json:"id" uri:"id" form:"id" binding:"required"`
	TenancyId uint `json:"tenancy_id" uri:"tenancy_id" form:"tenancy_id"`
	PatientId uint `json:"patient_id" uri:"patient_id" form:"patient_id"`
	UserId    uint `json:"user_id" uri:"user_id" form:"user_id"`
}

type IdsReq struct {
	Ids []uint `json:"ids" form:"ids" binding:"required"`
}

type DateReq struct {
	Date string `json:"date" form:"date"`
}

// Get role by id structure
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId" binding:"required"`
}

type Result struct {
	Time  string
	Count float64
}

type Empty struct{}
