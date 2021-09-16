package request

type BrandPageInfo struct {
	BrandCategoryId int32 `json:"brandCategoryId" form:"brandCategoryId"`
	PageInfo
}
