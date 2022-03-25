package front

import (
	"github.com/kataras/iris/v12"
	"server/model"
	"server/service"
)

type TagController struct {
	service *service.TagService
}

func (t *TagController) GetList() (result []*model.Tag, code int) {
	result, err := t.service.List()
	if err != nil {
		code = iris.StatusInternalServerError
	}
	return
}

func NewTagController(service *service.TagService) *TagController {
	return &TagController{service}
}
