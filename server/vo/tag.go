package vo

import "server/model"

type TagListResult struct {
	Result
}

func (r *TagListResult) DataOk(data *[]*model.Tag) {
	r.Data = data
}
