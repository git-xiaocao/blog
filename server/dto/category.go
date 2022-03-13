package dto

type CategoryAddDTO struct {
	Name string `json:"name"`
}

type CategoryDeleteDTO struct {
	Id int64 `json:"id"`
}
