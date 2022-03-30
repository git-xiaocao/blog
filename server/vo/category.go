package vo

import "server/model"

type CategoryListResult struct {
	Result
}

func (r *CategoryListResult) DataOk(data *[]*model.Category) {
	r.Data = data
}
