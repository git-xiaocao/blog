package dao

import "server/model"

type CategoryDAO struct{}

func (*CategoryDAO) QueryAll() (result []*model.Category, err error) {
	err = db.Model(&model.Category{}).Find(&result).Error
	return
}

func (*CategoryDAO) ExistByName(name string) (result bool) {
	db.Raw("SELECT TRUE FROM category WHERE name = ?", name).Scan(&result)
	return
}

func (*CategoryDAO) Add(category model.Category) (err error) {
	err = db.Save(&category).Error
	return
}

func NewCategoryDAO() *CategoryDAO {
	return &CategoryDAO{}
}
