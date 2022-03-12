package dao

import "server/model"

type UserDAO struct{}

func (*UserDAO) QueryList(offset int, limit int) (result []*model.User, err error) {
	err = db.Model(&model.User{}).
		Offset(offset).Limit(limit).
		Find(&result).Error
	return
}

func (*UserDAO) QueryById(id int64) (result model.User, err error) {
	err = db.Model(&model.User{}).
		Where("id = ?", id).
		Last(&result).Error
	return
}

func (*UserDAO) Add(user model.User) (err error) {
	err = db.Save(&user).Error
	return
}

func NewUserDAO() *UserDAO {
	return &UserDAO{}
}
