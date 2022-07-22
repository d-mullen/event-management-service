/*
 * Zenoss CONFIDENTIAL
 * __________________
 *
 *  This software Copyright (c) Zenoss, Inc. 2019
 *  All Rights Reserved.
 *
 * NOTICE:  All information contained herein is, and remains the property of Zenoss Incorporated
 * and its suppliers, if any.  The intellectual and technical concepts contained herein are owned
 * and proprietary to Zenoss Incorporated and its suppliers and may be covered by U.S. and Foreign
 * Patents, patents in process, and are protected by U.S. and foreign trade secret or copyright law.
 * Dissemination of this information or reproduction of any this material herein is strictly forbidden
 * unless prior written permission by an authorized officer is obtained from Zenoss Incorporated.
 */
package frequency

import (
	"bytes"
	"io"

	"math"

	"github.com/zenoss/zingo/v4/rowkey"
)

type Values interface {
	GetValues() []any
	String() string
	Bytes() []byte
}

type values string

func ValuesFromBytes(b []byte) Values {
	return values(string(b))
}

func (vs values) GetValues() (values []any) {
	dec := rowkey.NewRowKeyDecoderFromString(string(vs))
	defer dec.Free()

	for {
		v, err := dec.ReadUntyped()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if vint, ok := v.(int8); ok {
			v = vint == math.MaxInt8
		}
		values = append(values, v)
	}
	return
}

func (vs values) String() string {
	return string(vs)
}

func (vs values) Bytes() []byte {
	return []byte(vs)
}

func ToValues(vals ...any) Values {
	b := new(bytes.Buffer)
	for _, value := range vals {
		b.Write(mustMarshal(value))
	}
	return values(b.String())
}
