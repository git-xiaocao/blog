package model

type User struct {
	BaseModel
	Nick     string `json:"nick"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
