package vo

type EmptyResult struct {
	Result
}

func (r *EmptyResult) Ok() {
}
