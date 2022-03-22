package backstage

import (
	"server/service"
	"server/vo"
)

type CategoryController struct {
	service *service.CategoryService
}

func (c *CategoryController) GetList() (result *vo.Result) {
	data, err := c.service.List()

	if err != nil {
		result = &vo.Result{Error: true, Message: err.Error()}
	} else {
		result = &vo.Result{Data: data}
	}
	return
}

func NewCategoryController(service *service.CategoryService) *CategoryController {
	return &CategoryController{service}
}
