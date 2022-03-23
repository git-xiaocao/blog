package backstage

import (
	"github.com/kataras/iris/v12"
	"server/dto"
	"server/service"
	"server/vo"
)

type TagController struct {
	service *service.TagService
}

func (t *TagController) PostAdd(ctx iris.Context) (result vo.EmptyResult) {
	var body dto.TagAddDTO
	err := ctx.ReadJSON(&body)
	if err != nil {
		result.ReadJsonError(err)
		return
	}

	err = t.service.Add(body)
	if err != nil {
		result.DBError(err)
		return
	}
	result.Ok()
	return
}

func (t *TagController) GetList() (result vo.TagListResult) {
	data, err := t.service.List()
	if err != nil {
		result.DBError(err)
		return
	}
	result.DataOk(&data)

	return
}

func NewTagController(service *service.TagService) *TagController {
	return &TagController{service}
}
