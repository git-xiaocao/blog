package backstage

import (
	"github.com/kataras/iris/v12"
	"server/dto"
	"server/service"
	"server/util"
)

type ArticleController struct {
	service *service.ArticleService
}

//GetBy /article/{pathMark}
func (a *ArticleController) GetBy(pathMark string) (result *dto.ResultDTO, code int) {
	if a.service.Exist(pathMark) {
		data, err := a.service.Get(pathMark)
		if err != nil {
			result = &dto.ResultDTO{
				Error:   true,
				Message: err.Error(),
			}

		} else {
			result = &dto.ResultDTO{
				Data: data,
			}

		}
	} else {
		code = iris.StatusNotFound
	}
	return
}

//PostAdd /article/add
func (a *ArticleController) PostAdd(ctx iris.Context) (result *dto.ResultDTO, err error) {
	var body dto.ArticleAddDTO
	err = ctx.ReadJSON(&body)

	if err != nil {
		return
	}

	if e := a.service.Add(body); e != nil {
		result = &dto.ResultDTO{
			Error: true,
			Data:  e.Error(),
		}
	} else {
		result = &dto.ResultDTO{}
	}

	return
}

//PostUpdate /article/update
func (a *ArticleController) PostUpdate(ctx iris.Context) (result *dto.ResultDTO, err error) {
	var body dto.ArticleUpdateDTO
	err = ctx.ReadJSON(&body)

	if err != nil {
		return
	}
	if e := a.service.Update(body); e != nil {
		result = &dto.ResultDTO{
			Error:   true,
			Message: e.Error(),
		}
	} else {
		result = &dto.ResultDTO{}
	}

	return
}

//GetList /article/list
func (a *ArticleController) GetList(ctx iris.Context) (result *dto.ResultDTO, err error) {

	tagIds, err := util.QueryParamsInt64Array("tag-ids", ctx.URLParams())
	if err != nil {
		return
	}

	result = &dto.ResultDTO{Data: tagIds}
	return
}

func NewArticleController(service *service.ArticleService) *ArticleController {
	return &ArticleController{service}
}
