package dao

import (
	"gorm.io/gorm"
	"server/model"
)

type ArticleDAO struct{}

func (*ArticleDAO) QueryList(offset int, limit int) (result []*model.Article, err error) {

	err = db.Model(&model.Article{}).
		Preload("CategoryDAO").Preload("Tags").
		Offset(offset).Limit(limit).
		Find(&result).Error

	return
}

func (*ArticleDAO) QueryById(id int64) (result model.Article, err error) {
	err = db.Model(&model.Article{}).
		Where("id = ?", id).
		Preload("CategoryDAO").Preload("Tags").Last(&result).Error
	return
}

func (*ArticleDAO) QueryByTagIds(ids []int64, offset int, limit int) (result []*model.Article, err error) {

	categoryIds := make([]int, 0)
	err = db.Transaction(func(tx *gorm.DB) error {
		err := tx.Raw("SELECT article_id FROM article_tag WHERE tag_id IN ?", ids).Scan(&categoryIds).Error
		err = tx.Model(&model.Article{}).
			Preload("CategoryDAO").Preload("Tags").
			Where("id IN ?", categoryIds).
			Offset(offset).Limit(limit).
			Find(&result).Error
		return err
	})

	return
}

func (*ArticleDAO) Add() (err error) {
	err = db.Save(&model.Article{}).Error
	return
}

func (*ArticleDAO) Update(article model.Article) (err error) {
	err = db.Transaction(func(tx *gorm.DB) error {
		//清空Tag
		err := tx.Exec("DELETE FROM article_tag WHERE article_id = ?", article.Id).Error
		err = tx.Model(&article).UpdateColumns(&article).Error
		return err
	})

	return
}

func (*ArticleDAO) ExistByPathMark(pathMark string) (result bool) {
	db.Raw("SELECT TRUE FROM article WHERE name = ?", pathMark).Scan(&result)
	return
}

func NewArticleDAO() *ArticleDAO {
	return &ArticleDAO{}
}
