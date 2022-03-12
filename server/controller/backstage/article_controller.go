package backstage

import (
	"github.com/kataras/iris/v12/mvc"
	"server/service"
)

type ArticleController struct {
	service *service.ArticleService
}

func (*ArticleController) BeforeActivation(mvc.BeforeActivation) {

	//b.Handle("GET", "/{pathMark:string}", "GetByPathMark")
}

func (*ArticleController) GetBy(pathMark string) string {

	return "pathMark:" + pathMark
}

func NewArticleController(service *service.ArticleService) *ArticleController {
	return &ArticleController{service}
}
