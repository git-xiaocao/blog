package model

type Article struct {
	BaseModel
	CategoryId int64    `json:"categoryId"`
	Category   Category `json:"category"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	PathMark   string   `json:"pathMark"`
	Tags       []*Tag   `json:"tags" gorm:"many2many:article_tag;"`
}
