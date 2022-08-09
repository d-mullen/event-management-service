package event

import (
	"encoding/json"
	"fmt"
)

func unwrapFilterValue(v any) (any, error) {
	if vMap, ok := v.(map[string]any); ok {
		return filterFromAnyMap(vMap)
	}
	return v, nil
}

func filterFromAnyMap(fMap map[string]any) (*Filter, error) {
	var (
		f           Filter
		field       string
		op          FilterOp
		filterValue any
	)
	if fieldAny, ok := fMap["field"].(string); ok {
		field = fieldAny
	} else {
		return nil, fmt.Errorf("unmarshal failed: invalid type of field: expected string, got %T", fMap["field"])
	}
	if opAny, ok := fMap["op"].(string); ok {
		op = FilterOp(opAny)
	} else {
		return nil, fmt.Errorf("unmarshal failed: invalid type of field: expected %T, got %T", FilterOpAnd, fMap["field"])
	}
	f.Field = field
	f.Op = op
	ok := false
	filterValue, ok = fMap["value"]
	if !ok {
		return nil, fmt.Errorf("uh oh")
	}
	switch otherFilterV := filterValue.(type) {
	case map[string]any:
		otherFilter, err := filterFromAnyMap(otherFilterV)
		if err != nil {
			return nil, err
		}
		f.Value = otherFilter
	case []any:
		for i, otherValue := range otherFilterV {
			otherFilter, err := unwrapFilterValue(otherValue)
			if err != nil {
				return nil, err
			}
			otherFilterV[i] = otherFilter
		}
		f.Value = otherFilterV
	default:
		v, err := unwrapFilterValue(filterValue)
		if err != nil {
			return nil, err
		}
		f.Value = v
	}
	return &f, nil
}

func (f *Filter) UnmarshalJSON(b []byte) error {
	fMap := make(map[string]any)
	err := json.Unmarshal(b, &fMap)
	if err != nil {
		return err
	}
	_f, err := filterFromAnyMap(fMap)
	if err != nil {
		return err
	}
	f.Field = _f.Field
	f.Op = _f.Op
	f.Value = _f.Value
	return nil
}
