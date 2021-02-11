/*
 * Zenoss CONFIDENTIAL
 * __________________
 *
 *  This software Copyright (c) Zenoss, Inc. 2018
 *  All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and remains the property of Zenoss Incorporated
 * and its suppliers, if any.  The intellectual and technical concepts contained herein are owned
 * and proprietary to Zenoss Incorporated and its suppliers and may be covered by U.S. and Foreign
 * Patents, patents in process, and are protected by U.S. and foreign trade secret or copyright law.
 * Dissemination of this information or reproduction of any this material herein is strictly forbidden
 * unless prior written permission by an authorized officer is obtained from Zenoss Incorporated.
 */

package protobufutils

import (
	pbstruct "github.com/golang/protobuf/ptypes/struct"
	"github.com/pkg/errors"

	"github.com/zenoss/zing-proto/v11/go/cloud/common"
)

var (
	ErrInvalidInput = errors.New("can't convert input from/to protobuf struct")
)

// ToProtoListValue builds an slice of scalars from the passed input
// It errors if the input data contains non scalar values
func ToProtoListValue(vv []interface{}) (*pbstruct.Value, error) {
	var values []*pbstruct.Value = make([]*pbstruct.Value, 0, len(vv))
	for _, v := range vv {
		var value *pbstruct.Value
		switch x := v.(type) {
		case string:
			value = &pbstruct.Value{Kind: &pbstruct.Value_StringValue{StringValue: x}}
		case int:
			value = &pbstruct.Value{Kind: &pbstruct.Value_NumberValue{NumberValue: float64(x)}}
		case int8:
			value = &pbstruct.Value{Kind: &pbstruct.Value_NumberValue{NumberValue: float64(x)}}
		case int16:
			value = &pbstruct.Value{Kind: &pbstruct.Value_NumberValue{NumberValue: float64(x)}}
		case int32:
			value = &pbstruct.Value{Kind: &pbstruct.Value_NumberValue{NumberValue: float64(x)}}
		case int64:
			value = &pbstruct.Value{Kind: &pbstruct.Value_NumberValue{NumberValue: float64(x)}}
		case uint:
			value = &pbstruct.Value{Kind: &pbstruct.Value_NumberValue{NumberValue: float64(x)}}
		case uint8:
			value = &pbstruct.Value{Kind: &pbstruct.Value_NumberValue{NumberValue: float64(x)}}
		case uint16:
			value = &pbstruct.Value{Kind: &pbstruct.Value_NumberValue{NumberValue: float64(x)}}
		case uint32:
			value = &pbstruct.Value{Kind: &pbstruct.Value_NumberValue{NumberValue: float64(x)}}
		case uint64:
			value = &pbstruct.Value{Kind: &pbstruct.Value_NumberValue{NumberValue: float64(x)}}
		case float32:
			value = &pbstruct.Value{Kind: &pbstruct.Value_NumberValue{NumberValue: float64(x)}}
		case float64:
			value = &pbstruct.Value{Kind: &pbstruct.Value_NumberValue{NumberValue: x}}
		case bool:
			value = &pbstruct.Value{Kind: &pbstruct.Value_BoolValue{BoolValue: x}}
		default:
			msg := "unexpected type %T. proto struct field values must be scalar"
			return nil, errors.Wrapf(ErrInvalidInput, msg, x)
		}
		if value != nil {
			values = append(values, value)
		}
	}
	return &pbstruct.Value{Kind: &pbstruct.Value_ListValue{ListValue: &pbstruct.ListValue{Values: values}}}, nil
}

// ToProtoStruct converts the received input into a protobuf struct. The input map values
// must be an slice of scalars or else it will return an error
func ToProtoStruct(vv map[string][]interface{}) (*pbstruct.Struct, error) {
	fields := map[string]*pbstruct.Value{}
	for k, v := range vv {
		pbv, err := ToProtoListValue(v)
		if err != nil {
			return nil, err
		}
		fields[k] = pbv
	}
	return &pbstruct.Struct{Fields: fields}, nil
}

func unexpectedFieldType(v interface{}) error {
	return errors.Wrapf(ErrInvalidInput, "unexpected field type %T / %+v", v, v)
}

func listValueToSliceOfInterface(vv *pbstruct.ListValue) ([]interface{}, error) {
	values := make([]interface{}, 0, len(vv.Values))
	for _, v := range vv.Values {
		switch x := v.GetKind().(type) {
		case *pbstruct.Value_BoolValue:
			values = append(values, x.BoolValue)
		case *pbstruct.Value_NumberValue:
			values = append(values, x.NumberValue)
		case *pbstruct.Value_StringValue:
			values = append(values, x.StringValue)
		default:
			return nil, unexpectedFieldType(x)
		}
	}
	return values, nil
}

func StructToMapOfScalarArray(s *pbstruct.Struct) (map[string]*common.ScalarArray, error) {
	var err error
	out := map[string]*common.ScalarArray{}
	for fName, fValue := range s.GetFields() {
		outValue := []interface{}{}
		switch x := fValue.GetKind().(type) {
		case *pbstruct.Value_BoolValue:
			outValue = append(outValue, x.BoolValue)
		case *pbstruct.Value_NumberValue:
			outValue = append(outValue, x.NumberValue)
		case *pbstruct.Value_StringValue:
			outValue = append(outValue, x.StringValue)
		case *pbstruct.Value_ListValue:
			outValue, err = listValueToSliceOfInterface(x.ListValue)
			if err != nil {
				return nil, err
			}
		default:
			return nil, unexpectedFieldType(x)
		}
		out[fName], err = ToScalarArray(outValue)
		if err != nil {
			return nil, errors.Wrap(err, "error converting field value to ScalarArray")
		}
	}
	return out, nil
}
