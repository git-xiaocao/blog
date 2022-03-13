package dao

import (
	"server/model"
)

type TagDAO struct{}

func (*TagDAO) QueryAll() (result []*model.Tag, err error) {
	err = db.Model(&model.Tag{}).Find(&result).Error
	return
}

func (*TagDAO) ExistByName(name string) (result bool) {
	db.Raw("SELECT TRUE FROM tag WHERE name = ?", name).Scan(&result)
	return
}

func (*TagDAO) Add(tag model.Tag) (err error) {
	err = db.Save(&tag).Error
	return
}

func (*TagDAO) DeleteById(id int64) (err error) {
	err = db.Delete(&model.Tag{}).Where("id = ?", id).Error
	return
}

func NewTagDAO() *TagDAO {
	return &TagDAO{}
}
