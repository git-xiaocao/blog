package dto

//UserRegisterDTO 用户注册
type UserRegisterDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

//UserChangePasswordDTO 用户修改密码
type UserChangePasswordDTO struct {
	Email       string `json:"email"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

//UserRetrievePassword 用户找回密码
type UserRetrievePassword struct {
	Email       string `json:"email"`
	Code        string `json:"code"`
	NewPassword string `json:"newPassword"`
}
