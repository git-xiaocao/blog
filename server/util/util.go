package util

import (
	"reflect"
	"strconv"
	"strings"
)

func QueryParamsArray[T int8 | int16 | int32 | int64 | string](key string, m map[string]string) (result []T, err error) {
	var t T
	typeName := reflect.TypeOf(t).Name()
	result = make([]T, 0)
	for k, v := range m {
		if k == key {
			arr := strings.Split(v, ",")

			for _, s := range arr {
				var temp interface{}
				switch typeName {
				case "int8":
					temp, err = strconv.ParseInt(s, 10, 8)
					break
				case "int16":
					temp, err = strconv.ParseInt(s, 10, 16)
					break
				case "int32":
					temp, err = strconv.ParseInt(s, 10, 32)
					break
				case "int64":
					temp, err = strconv.ParseInt(s, 10, 64)
					break
				case "string":
					temp = s
					break
				}

				if err != nil {
					result = nil
					return
				}
				result = append(result, temp.(T))

			}

		}

	}
	return
}
