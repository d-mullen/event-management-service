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
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/zenoss/zing-proto/v11/go/cloud/common"
)

var (
	ErrInvalidInput = errors.New("can't convert input from/to protobuf struct")
)

// ToProtoListValue builds an slice of scalars from the passed input
// It errors if the input data contains non scalar values
func ToProtoListValue(vv []interface{}) (*structpb.Value, error) {
	var values []*structpb.Value = make([]*structpb.Value, 0, len(vv))
	for _, v := range vv {
		var value *structpb.Value
		switch x := v.(type) {
		case string:
			value = &structpb.Value{Kind: &structpb.Value_StringValue{StringValue: x}}
		case int:
			value = &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}
		case int8:
			value = &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}
		case int16:
			value = &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}
		case int32:
			value = &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}
		case int64:
			value = &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}
		case uint:
			value = &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}
		case uint8:
			value = &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}
		case uint16:
			value = &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}
		case uint32:
			value = &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}
		case uint64:
			value = &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}
		case float32:
			value = &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(x)}}
		case float64:
			value = &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: x}}
		case bool:
			value = &structpb.Value{Kind: &structpb.Value_BoolValue{BoolValue: x}}
		default:
			msg := "unexpected type %T. proto struct field values must be scalar"
			return nil, errors.Wrapf(ErrInvalidInput, msg, x)
		}
		if value != nil {
			values = append(values, value)
		}
	}
	return &structpb.Value{Kind: &structpb.Value_ListValue{ListValue: &structpb.ListValue{Values: values}}}, nil
}

// ToProtoStruct converts the received input into a protobuf struct. The input map values
// must be an slice of scalars or else it will return an error
func ToProtoStruct(vv map[string][]interface{}) (*structpb.Struct, error) {
	fields := map[string]*structpb.Value{}
	for k, v := range vv {
		pbv, err := ToProtoListValue(v)
		if err != nil {
			return nil, err
		}
		fields[k] = pbv
	}
	return &structpb.Struct{Fields: fields}, nil
}

func unexpectedFieldType(v interface{}) error {
	return errors.Wrapf(ErrInvalidInput, "unexpected field type %T / %+v", v, v)
}

func listValueToSliceOfInterface(vv *structpb.ListValue) ([]interface{}, error) {
	values := make([]interface{}, 0, len(vv.Values))
	for _, v := range vv.Values {
		switch x := v.GetKind().(type) {
		case *structpb.Value_BoolValue:
			values = append(values, x.BoolValue)
		case *structpb.Value_NumberValue:
			values = append(values, x.NumberValue)
		case *structpb.Value_StringValue:
			values = append(values, x.StringValue)
		default:
			return nil, unexpectedFieldType(x)
		}
	}
	return values, nil
}

func StructToMapOfScalarArray(s *structpb.Struct) (map[string]*common.ScalarArray, error) {
	var err error
	out := map[string]*common.ScalarArray{}
	for fName, fValue := range s.GetFields() {
		outValue := []interface{}{}
		switch x := fValue.GetKind().(type) {
		case *structpb.Value_BoolValue:
			outValue = append(outValue, x.BoolValue)
		case *structpb.Value_NumberValue:
			outValue = append(outValue, x.NumberValue)
		case *structpb.Value_StringValue:
			outValue = append(outValue, x.StringValue)
		case *structpb.Value_ListValue:
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
