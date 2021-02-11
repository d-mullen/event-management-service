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

import "github.com/golang/protobuf/proto"

type ProtobufMarshaller interface {
	Marshal(proto.Message) ([]byte, error)
	Unmarshal([]byte, proto.Message) error
}

func NewProtobufMarshaller() ProtobufMarshaller {
	return &protobufMarshaller{}
}

type protobufMarshaller struct{}

func (m *protobufMarshaller) Marshal(pb proto.Message) ([]byte, error) {
	return proto.Marshal(pb)
}

func (m *protobufMarshaller) Unmarshal(b []byte, pb proto.Message) error {
	return proto.Unmarshal(b, pb)
}
