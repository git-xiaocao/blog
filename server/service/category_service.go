package service

import (
	"server/dao"
	"server/dto"
	"server/model"
)

type CategoryService struct {
	dao *dao.CategoryDAO
}

func (c *CategoryService) Add(dto dto.CategoryAddDTO) (err error) {
	category := model.Category{Name: dto.Name}
	err = c.dao.Add(category)
	return
}

func (c *CategoryService) Delete(dto dto.CategoryDeleteDTO) (err error) {
	err = c.dao.DeleteById(dto.Id)
	return
}

func (c *CategoryService) List() (result []*model.Category, err error) {
	result, err = c.dao.QueryAll()
	return
}

func NewCategoryService(dao *dao.CategoryDAO) *CategoryService {
	return &CategoryService{dao}
}
