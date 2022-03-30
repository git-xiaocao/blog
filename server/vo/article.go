package vo

import "server/model"

type ArticleResult struct {
	Result
}

func (r *ArticleResult) DataOk(data *model.Article) {
	r.Data = data
}

type ArticleListResult struct {
	Result
}

func (r *ArticleListResult) DataOk(data *[]*model.Article) {
	r.Data = data
}

type ArticlePageResult struct {
	PageResult[*[]*model.Article]
}

func (r *ArticlePageResult) DataOk(data *PageData[[]*model.Article]) {
	r.Data = data
}
