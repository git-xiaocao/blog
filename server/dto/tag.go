package dto

type TagAddDTO struct {
	Name string `json:"name"`
}

type TagDeleteDTO struct {
	Id int64 `json:"id"`
}
