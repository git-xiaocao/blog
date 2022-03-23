package vo

import "server/model"

type ArticleResult struct {
	Result[*model.Article]
}

func (r *ArticleResult) DataOk(data *model.Article) {
	var temp interface{}
	temp = data
	r.Data = temp
}

type ArticleListResult struct {
	Result[*[]*model.Article]
}

func (r *ArticleListResult) DataOk(data *[]*model.Article) {
	var temp interface{}
	temp = data
	r.Data = temp
}
