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

func (*CategoryDAO) ExistById(id int64) (result bool) {
	db.Raw("SELECT TRUE FROM category WHERE id = ?", id).Scan(&result)
	return
}

func (*CategoryDAO) Add(category model.Category) (err error) {
	err = db.Save(&category).Error
	return
}

func (*CategoryDAO) DeleteById(id int64) (err error) {
	err = db.Delete(&model.Category{}).Where("id = ?", id).Error
	return
}

func NewCategoryDAO() *CategoryDAO {
	return &CategoryDAO{}
}
