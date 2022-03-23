package vo

type EmptyResult struct {
	Result[*any]
}

func (r *EmptyResult) Ok() {
}
