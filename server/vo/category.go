package vo

import "server/model"

type CategoryListResult struct {
	Result[*[]*model.Category]
}

func (r *CategoryListResult) DataOk(data *[]*model.Category) {
	var temp interface{}
	temp = data
	r.Data = temp
}
