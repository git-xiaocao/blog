package model

type ArticleTag struct {
	Id        int64 `json:"id"`
	ArticleId int64 `json:"articleId"`
	TagId     int64 `json:"tagId"`
}
