/*
 * Zenoss CONFIDENTIAL
 * __________________
 *
 *  This software Copyright (c) Zenoss, Inc. 2020
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
	"encoding/json"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// ProtoMsgToMap converts proto.Message to the map using jsonpb Marshaler
func ProtoMsgToMap(data proto.Message) (map[string]interface{}, error) {
	var mapValue map[string]interface{}
	m := jsonpb.Marshaler{}
	jsonStr, err := m.MarshalToString(data)
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(jsonStr), &mapValue)
	return mapValue, nil
}
