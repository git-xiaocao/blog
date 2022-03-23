package service

import (
	"server/dao"
	"server/dto"
	"server/model"
)

type TagService struct {
	dao *dao.TagDAO
}

func (t *TagService) Add(dto dto.TagAddDTO) (err error) {
	tag := model.Tag{Name: dto.Name}
	err = t.dao.Add(tag)
	return
}

func (t *TagService) Delete(dto dto.TagDeleteDTO) (err error) {
	err = t.dao.DeleteById(dto.Id)
	return
}

func (t *TagService) List() (result []*model.Tag, err error) {
	result, err = t.dao.QueryAll()
	return
}

func NewTagService(dao *dao.TagDAO) *TagService {
	return &TagService{dao}
}
