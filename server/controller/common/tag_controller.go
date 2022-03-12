package common

import (
	"github.com/kataras/iris/v12"
	"server/model"
	"server/service"
)

type TagController struct {
	service *service.TagService
}

func (t *TagController) GetAll() (result []*model.Tag, code int) {
	result, err := t.service.QueryAll()
	if err != nil {
		code = iris.StatusInternalServerError
	}
	return
}

func NewTagController(service *service.TagService) *TagController {
	return &TagController{service}
}
