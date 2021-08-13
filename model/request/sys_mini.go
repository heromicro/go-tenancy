package request

type CreateSysMini struct {
	Name      string `json:"name" form:"name" binding:"required"`
	AppID     string `json:"appId" form:"appId" binding:"required"`
	AppSecret string `json:"appSecret" form:"appSecret" binding:"required"`
	Remark    string `json:"remark" form:"remark"`
}

type UpdateSysMini struct {
	Name      string `json:"name" form:"name" binding:"required"`
	AppID     string `json:"appId" form:"appId" binding:"required"`
	AppSecret string `json:"appSecret" form:"appSecret" binding:"required"`
	Remark    string `json:"remark" form:"remark"`
}
