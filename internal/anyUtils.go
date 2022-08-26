package internal

import (
	"reflect"
)

func AnyToSlice(v any) []any {
	var (
		out []any
	)
	rv := reflect.ValueOf(v)
	if k := rv.Kind(); k == reflect.Slice {
		for i := 0; i < rv.Len(); i++ {
			out = append(out, rv.Index(i).Interface())
		}
	} else {
		return []any{v}
	}
	return out
}

func CloneSlice[E any](arr []E) []E {
	if arr == nil {
		return nil
	}
	retVal := make([]E, len(arr))
	for i, v := range arr {
		retVal[i] = v
	}
	return retVal
}
