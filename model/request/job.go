package request

type Task struct {
	Name string `json:"name"  uri:"name" form:"name"  binding:"required"`
}
