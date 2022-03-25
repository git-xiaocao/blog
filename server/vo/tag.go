package vo

import "server/model"

type TagListResult struct {
	Result[*[]*model.Tag]
}

func (r *TagListResult) DataOk(data *[]*model.Tag) {
	r.Data = data
}
