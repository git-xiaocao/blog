package dao

import (
	"gorm.io/gorm"
	"server/model"
)

type ArticleDAO struct{}

func (*ArticleDAO) QueryList(offset, limit int) (result []*model.Article, err error) {

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

func (*ArticleDAO) QueryByPathMark(pathMark string) (result model.Article, err error) {
	err = db.Model(&model.Article{}).
		Where("path_mark = ?", pathMark).
		Preload("CategoryDAO").Preload("Tags").Last(&result).Error
	return
}

func (*ArticleDAO) QueryListByTagIds(tagIds []int64, offset, limit int) (result []*model.Article, err error) {

	articleIds := make([]int, 0)
	err = db.Transaction(func(tx *gorm.DB) error {
		//先把文章id查出来
		err := tx.Raw("SELECT article_id FROM article_tag WHERE tag_id IN ?", tagIds).Scan(&articleIds).Error
		if err != nil {
			return err
		}
		//再用文章id去查文章
		err = tx.Model(&model.Article{}).
			Preload("CategoryDAO").Preload("Tags").
			Where("id IN ?", articleIds).
			Offset(offset).Limit(limit).
			Find(&result).Error
		return err
	})

	return
}

func (*ArticleDAO) Add(article model.Article) (err error) {
	err = db.Save(&article).Error
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

func (*ArticleDAO) DeleteById(id int64) (err error) {
	err = db.Delete(&model.Article{}).Where("id = ?", id).Error
	return
}

func NewArticleDAO() *ArticleDAO {
	return &ArticleDAO{}
}
