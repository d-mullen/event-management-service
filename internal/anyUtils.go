package internal

import (
	"fmt"
	"reflect"
)

func AnyToSlice(v any) ([]any, error) {
	var (
		out []any
	)
	rv := reflect.ValueOf(v)
	if k := rv.Kind(); k == reflect.Slice {
		for i := 0; i < rv.Len(); i++ {
			out = append(out, rv.Index(i).Interface())
		}
	} else {
		return nil, fmt.Errorf("failed to convert any: %v to []any: got %v, expected %v", v, k, reflect.Slice)
	}
	return out, nil
}

func MustAnyToSlice(v any) []any {
	anySlice, err := AnyToSlice(v)
	if err != nil {
		panic(err)
	}
	return anySlice
}
