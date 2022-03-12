package common

import (
	"server/service"
)

type ArticleController struct {
	service *service.ArticleService
}

func (*ArticleController) GetBy(pathMark string) string {

	return "pathMark:" + pathMark
}

func NewArticleController(service *service.ArticleService) *ArticleController {
	return &ArticleController{service}
}
