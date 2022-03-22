package vo

import "fmt"

type Result[T any] struct {
	HasError bool   `json:"hasError"`
	Error    string `json:"error"`
	Message  string `json:"message"`
	Data     T      `json:"data"`
}

func (r *Result[T]) Err(err error, message string) {
	r.HasError = true
	r.Error = err.Error()
	r.Message = message
}

func (r *Result[T]) DataOk(data T) {
	r.Data = data
}

func (r *Result[T]) Ok() {
}

func (r *Result[T]) DBError(err error) {
	r.HasError = true
	r.Error = err.Error()
	r.Message = "数据库错误"
}

func (r *Result[T]) ReadJsonError(err error) {
	r.HasError = true
	r.Error = err.Error()
	r.Message = "序列化JSON错误"
}

func (r *Result[T]) ReadQueryError(err error) {
	r.HasError = true
	r.Error = err.Error()
	r.Message = "获取查询参数错误"
}

func (r *Result[T]) QueryParamLack(paramName string) {
	r.HasError = true
	r.Message = fmt.Sprintf("查询参数%s缺失", paramName)
}
