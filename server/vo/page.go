package vo

import "math"

type PageResult[T any] struct {
	Result
}

type PageData[T any] struct {
	PageCurrent int         `json:"pageCurrent"`
	PageTotal   int         `json:"pageTotal"`
	Total       int64       `json:"total"`
	List        interface{} `json:"list"`
}

func (p *PageData[T]) Init(list interface{}, pageCurrent, pageSize int, count int64) {
	p.PageCurrent = pageCurrent
	//数据总数 / 每页数量 向上取整
	p.PageTotal = int(math.Ceil(float64(count) / float64(pageSize)))
	p.Total = count
	p.List = list
}
