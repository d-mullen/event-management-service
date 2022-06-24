/**************************************************************************
 * Zenoss CONFIDENTIAL
 * __________________
 *
 *  This software Copyright (c) Zenoss, Inc. 2005-2021
 *  All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and remains the property of Zenoss Incorporated
 * and its suppliers, if any.  The intellectual and technical concepts contained herein are owned
 * and proprietary to Zenoss Incorporated and its suppliers and may be covered by U.S. and Foreign
 * Patents, patents in process, and are protected by U.S. and foreign trade secret or copyright law.
 * Dissemination of this information or reproduction of any this material herein is strictly forbidden
 * unless prior written permission by an authorized officer is obtained from Zenoss Incorporated.
************************************************************************/

package protobufutils

import (
	"context"

	pbstruct "github.com/golang/protobuf/ptypes/struct"
	"github.com/zenoss/zing-proto/v11/go/cloud/common"
	"google.golang.org/protobuf/types/known/structpb"
)

func StringToListValue(str string) *pbstruct.Value {
	return &pbstruct.Value{
		Kind: &pbstruct.Value_ListValue{
			ListValue: &pbstruct.ListValue{
				Values: []*pbstruct.Value{StringToValue(str)},
			},
		},
	}
}

func StringToValue(str string) *pbstruct.Value {
	return &pbstruct.Value{
		Kind: &pbstruct.Value_StringValue{
			StringValue: str,
		},
	}
}

// ListValueToString extracts first value from the list
func ListValueToString(ctx context.Context, val *pbstruct.Value) string {
	list := val.GetListValue()
	if list != nil && len(list.Values) > 0 {
		return list.Values[0].GetStringValue()
	}
	return ""
}

// ListValueToString extracts first value from the list
func ListValueToStringSlice(ctx context.Context, listVal *pbstruct.ListValue) []string {
	if listVal != nil && len(listVal.Values) > 0 {
		result := make([]string, 0, len(listVal.Values))
		for _, value := range listVal.Values {
			result = append(result, value.GetStringValue())
		}
		return result
	}
	return []string{}
}

func ScalarArrayToValue(scalarArray *common.ScalarArray) (*structpb.Value, error) {
	unpackedV, err := FromScalarArray(scalarArray)
	if err != nil {
		return nil, err
	}
	value, err := ToProtoListValue(unpackedV)
	if err != nil {
		return nil, err
	}
	return value, nil
}
