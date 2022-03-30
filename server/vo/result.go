package vo

import "fmt"

type Result struct {
	HasError bool        `json:"hasError"`
	Error    string      `json:"error"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}

func (r *Result) Err(err error, message string) {
	r.HasError = true
	if err != nil {
		r.Error = err.Error()
	}
	r.Message = message
}

func (r *Result) NotFound() {
	r.HasError = true
	r.Message = "找不到"
}

func (r *Result) DBError(err error) {
	r.HasError = true
	r.Error = err.Error()
	r.Message = "数据库错误"
}

func (r *Result) ReadJsonError(err error) {
	r.HasError = true
	r.Error = err.Error()
	r.Message = "序列化JSON错误"
}

func (r *Result) ReadQueryError(err error) {
	r.HasError = true
	r.Error = err.Error()
	r.Message = "获取查询参数错误"
}

func (r *Result) QueryParamLack(paramName string) {
	r.HasError = true
	r.Message = fmt.Sprintf("查询参数%s缺失", paramName)
}
