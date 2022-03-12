package model

import "gorm.io/gorm"

type Article struct {
	BaseModel
	CategoryId int64    `json:"categoryId"`
	Category   Category `json:"category"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	PathMark   string   `json:"pathMark"`
	Tags       []*Tag   `json:"tags" gorm:"many2many:article_tag;"`
}

func (a Article) Add() (err error) {
	err = db.Save(&a).Error
	return
}

func (a Article) Update() (err error) {
	err = db.Transaction(func(tx *gorm.DB) error {
		//清空Tag
		err := tx.Exec("DELETE FROM article_tag WHERE article_id = ?", a.Id).Error
		err = tx.Model(&a).UpdateColumns(&a).Error
		return err
	})

	return
}

func (a Article) QueryList(offset int, limit int) (err error, result []*Article) {

	err = db.Model(&a).
		Preload("Category").Preload("Tags").
		Offset(offset).Limit(limit).
		Find(&result).Error

	return
}

func (a Article) QueryById() (err error, result Article) {
	err = db.Model(&a).
		Where("id = ?", a.Id).
		Preload("Category").Preload("Tags").Last(&result).Error
	return
}

func (a Article) QueryByTagIds(ids []int64, offset int, limit int) (err error, result []*Article) {

	categoryIds := make([]int, 0)
	err = db.Transaction(func(tx *gorm.DB) error {
		err := tx.Raw("SELECT article_id FROM article_tag WHERE tag_id IN ?", ids).Scan(&categoryIds).Error
		err = tx.Model(&a).
			Preload("Category").Preload("Tags").
			Where("id IN ?", categoryIds).
			Offset(offset).Limit(limit).
			Find(&result).Error
		return err
	})

	return
}

func (a Article) ExistByPathMark() (result bool) {
	db.Raw("SELECT TRUE FROM article WHERE name = ?", a.PathMark).Scan(&result)
	return
}
