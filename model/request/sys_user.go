package request

// User register structure
type Register struct {
	Username        string   `json:"username" binding:"required"`
	Password        string   `json:"password" binding:"required"`
	ConfirmPassword string   `json:"confirmPassword" binding:"required"`
	AuthorityId     []string `json:"authorityId" binding:"required"`
	NickName        string   `json:"nickName"`
	Phone           string   `json:"phone"`
	Status          int      `json:"status"`
}

// User login structure
type Login struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Captcha   string `json:"captcha" binding:"dev-required"`
	CaptchaId string `json:"captchaId" binding:"dev-required"`
}

// Modify password structure
type ChangePassword struct {
	Password        string `json:"password" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

// Modify password structure
type ChangeProfile struct {
	NickName string `json:"nickName" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

// Modify  user's auth structure
type SetUserAuth struct {
	Id          uint   `json:"id" form:"id" binding:"required,gt=0"`
	AuthorityId string `json:"authorityId" binding:"required"`
}

// Modify  user's auth structure
type UpdateUser struct {
	Status      int      `json:"status"`
	Email       string   `json:"email"`
	Phone       string   `json:"phone"`
	NickName    string   `json:"nickName"`
	Username    string   `json:"username"  binding:"required"`
	HeaderImg   string   `json:"headerImg" `
	AuthorityId []string `json:"authorityId" binding:"required"`
}
