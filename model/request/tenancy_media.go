package request

type UpdateMediaName struct {
	Name string `json:"name" form:"name" binding:"required"`
}

type MediaPageInfo struct {
	PageInfo
	Name string `json:"name" form:"name"`
}
