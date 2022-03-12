package backstage

import "server/service"

type CategoryController struct {
	service *service.CategoryService
}

func NewCategoryController(service *service.CategoryService) *CategoryController {
	return &CategoryController{service}
}
