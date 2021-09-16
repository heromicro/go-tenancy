package request

// Paging common input parameter structure
type ExpressPageInfo struct {
	Name string `json:"name" form:"name"`
	PageInfo
}

type GetByCode struct {
	Code string `json:"code" uri:"code" form:"code"`
}
