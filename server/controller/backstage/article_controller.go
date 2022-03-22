package backstage

import (
	"github.com/kataras/iris/v12"
	"server/dto"
	"server/service"
	"server/util"
	"server/vo"
)

type ArticleController struct {
	service *service.ArticleService
}

//GetBy /article/{pathMark}
func (a *ArticleController) GetBy(pathMark string) (result *vo.Result, code int) {
	if a.service.Exist(pathMark) {
		data, err := a.service.Get(pathMark)
		if err != nil {
			result = &vo.Result{
				Error:   true,
				Message: err.Error(),
			}

		} else {
			result = &vo.Result{
				Data: data,
			}

		}
	} else {
		code = iris.StatusNotFound
	}
	return
}

//PostAdd /article/add
func (a *ArticleController) PostAdd(ctx iris.Context) (result *vo.Result, err error) {
	var body dto.ArticleAddDTO
	err = ctx.ReadJSON(&body)

	if err != nil {
		return
	}

	if e := a.service.Add(body); e != nil {
		result = &vo.Result{
			Error: true,
			Data:  e.Error(),
		}
	} else {
		result = &vo.Result{}
	}

	return
}

//PostUpdate /article/update
func (a *ArticleController) PostUpdate(ctx iris.Context) (result *vo.Result, err error) {
	var body dto.ArticleUpdateDTO
	err = ctx.ReadJSON(&body)

	if err != nil {
		return
	}
	if e := a.service.Update(body); e != nil {
		result = &vo.Result{
			Error:   true,
			Message: e.Error(),
		}
	} else {
		result = &vo.Result{}
	}

	return
}

//GetList /article/list
func (a *ArticleController) GetList(ctx iris.Context) (result *vo.Result, err error) {

	tagIds, err := util.QueryParamsArray[int64]("tag-ids", ctx.URLParams())
	if err != nil {
		return
	}

	result = &vo.Result{Data: tagIds}
	return
}

func NewArticleController(service *service.ArticleService) *ArticleController {
	return &ArticleController{service}
}
