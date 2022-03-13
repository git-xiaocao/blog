package util

import "strconv"

func QueryParamsInt64Array(key string, m map[string]string) (result []int64, err error) {
	result = make([]int64, 0)
	for k, v := range m {
		if k == key {
			i64, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				result = nil
				return
			}
			result = append(result, i64)
		}

	}
	return
}
