package vo

import "server/model"

type TagListResult struct {
	Result[*[]*model.Tag]
}

func (r *TagListResult) DataOk(data *[]*model.Tag) {
	var temp interface{}
	temp = data
	r.Data = temp
}
