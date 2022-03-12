package service

import "server/dao"

type CategoryService struct {
	dao *dao.CategoryDAO
}

func NewCategoryService(dao *dao.CategoryDAO) *CategoryService {
	return &CategoryService{dao}
}
