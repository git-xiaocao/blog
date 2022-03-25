package backstage

import (
	"github.com/kataras/iris/v12"
	"server/dto"
	"server/service"
	"server/vo"
)

type ArticleController struct {
	service *service.ArticleService
}

//GetBy /article/{pathMark}
func (a *ArticleController) GetBy(pathMark string) (result vo.ArticleResult, code int) {
	if !a.service.Exist(pathMark) {
		result.NotFound()
		code = iris.StatusNotFound
		return
	}
	data, err := a.service.Get(pathMark)
	if err != nil {
		result.DBError(err)
		return
	}

	result.DataOk(&data)
	return
}

//GetList /article/list
func (a *ArticleController) GetList(ctx iris.Context) (result vo.ArticleListResult) {
	var query struct {
		Offset *int    `url:"offset"`
		Limit  *int    `url:"limit"`
		TagIds []int64 `url:"tag-ids"`
	}
	err := ctx.ReadQuery(&query)
	if err != nil {
		result.ReadQueryError(err)
		return
	}

	if query.Offset == nil {
		result.QueryParamLack("offset")
		return
	}

	if query.Limit == nil {
		result.QueryParamLack("limit")
		return
	}

	//大于0小于30
	if !(*query.Limit >= 0 && *query.Limit <= 30) {
		result.Err(nil, "参数范围 0 <= limit >= 30")
		return
	}

	data, err := a.service.List(query.TagIds, *query.Offset, *query.Limit)

	result.DataOk(&data)
	return
}

//PostAdd /article/add
func (a *ArticleController) PostAdd(ctx iris.Context) (result vo.EmptyResult) {
	var body dto.ArticleAddDTO
	err := ctx.ReadJSON(&body)

	if err != nil {
		result.ReadJsonError(err)
		return
	}

	if err = a.service.Add(body); err != nil {
		result.Ok()
	} else {
		result.DBError(err)
	}

	return
}

//PostUpdate /article/update
func (a *ArticleController) PostUpdate(ctx iris.Context) (result vo.EmptyResult) {
	var body dto.ArticleUpdateDTO
	err := ctx.ReadJSON(&body)

	if err != nil {
		result.ReadJsonError(err)
		return
	}
	if err = a.service.Update(body); err != nil {
		result.DBError(err)
	} else {
		result.Ok()
	}

	return
}

func NewArticleController(service *service.ArticleService) *ArticleController {
	return &ArticleController{service}
}
