package service

import (
	"server/dao"
	"server/model"
)

type TagService struct {
	dao *dao.TagDAO
}

func (t *TagService) QueryAll() (result []*model.Tag, err error) {
	result, err = t.dao.QueryAll()
	return
}

func NewTagService(dao *dao.TagDAO) *TagService {
	return &TagService{dao}
}
