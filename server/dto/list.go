package dto

type PageListQuery struct {
	PageCurrent *int `url:"page-current"`
	PageSize    *int `url:"page-size"`
}

func (q *PageListQuery) Check() (ok bool, message string) {
	if q.PageCurrent == nil {
		message = "查询参数page-current缺失"
		return
	}
	if q.PageSize == nil {
		message = "查询参数page-size缺失"
		return
	}
	if !(*q.PageSize >= 0 && *q.PageSize <= 30) {
		message = "参数范围 0 <= page-size >= 30"
		return
	}
	ok = true
	return
}

type ListQuery struct {
	Offset *int `url:"offset"`
	Limit  *int `url:"limit"`
}

func (q *ListQuery) Check() (ok bool, message string) {
	if q.Offset == nil {
		message = "查询参数offset缺失"
		return
	}
	if q.Limit == nil {
		message = "查询参数limit缺失"
		return
	}

	if !(*q.Limit >= 0 && *q.Limit <= 30) {
		message = "参数范围 0 <= limit >= 30"
		return
	}

	ok = true
	return
}
