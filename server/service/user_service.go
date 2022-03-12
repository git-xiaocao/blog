package service

import "server/dao"

type UserService struct {
	dao *dao.UserDAO
}

func NewUserService(dao *dao.UserDAO) *UserService {
	return &UserService{dao}
}
