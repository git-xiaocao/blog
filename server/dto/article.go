package dto

type ArticleAddDTO struct {
	CategoryId int64    `json:"categoryId"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	PathMark   string   `json:"pathMark"`
	TagIds     []int64  `json:"tagIds"`
	NewTags    []string `json:"newTags"`
}

type ArticleUpdateDTO struct {
	Id int64 `json:"id"`
	ArticleAddDTO
}

type ArticleDeleteDTO struct {
	Id int64 `json:"id"`
}
