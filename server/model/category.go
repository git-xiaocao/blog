package model

type Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (c Category) Query() (err error, result []*Category) {
	err = db.Model(&c).Find(&result).Error
	return
}

func (c Category) ExistByName() (result bool) {
	db.Raw("SELECT TRUE FROM category WHERE name = ?", c.Name).Scan(&result)
	return
}

func (c Category) Add() (err error) {
	err = db.Save(&c).Error
	return
}
