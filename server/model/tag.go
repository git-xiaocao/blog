package model

type Tag struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (t Tag) Query() (err error, result []*Tag) {
	err = db.Model(&t).Find(&result).Error
	return
}

func (t Tag) ExistByName() (result bool) {
	db.Raw("SELECT TRUE FROM tag WHERE name = ?", t.Name).Scan(&result)
	return
}

func (t Tag) Add() (err error) {
	err = db.Save(&t).Error
	return
}
