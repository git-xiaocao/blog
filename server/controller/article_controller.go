package controller

import "github.com/kataras/iris/v12/mvc"

type ArticleController struct {
}

func (a *ArticleController) BeforeActivation(b mvc.BeforeActivation) {

	//b.Handle("GET", "/{pathMark:string}", "GetByPathMark")
}

func (a ArticleController) GetBy(pathMark string) string {
	return "pathMark:" + pathMark
}
