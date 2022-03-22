package backstage

import (
	"github.com/kataras/iris/v12"
	"server/dto"
	"server/model"
	"server/service"
	"server/vo"
)

type CategoryController struct {
	service *service.CategoryService
}

func (c *CategoryController) PostAdd(ctx iris.Context) (result vo.Result[*any]) {
	var body dto.CategoryAddDTO
	err := ctx.ReadJSON(&body)
	if err != nil {
		result.ReadJsonError(err)
		return
	}

	err = c.service.Add(body)
	if err != nil {
		result.DBError(err)
		return
	}
	result.Ok()

	return
}

func (c *CategoryController) GetList() (result vo.Result[*[]*model.Category]) {
	data, err := c.service.List()

	if err != nil {
		result.DBError(err)
		return
	}
	result.DataOk(&data)

	return
}

func NewCategoryController(service *service.CategoryService) *CategoryController {
	return &CategoryController{service}
}
