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
	"math"
	"time"

	"github.com/zenoss/zingo/v4/orderedbytes"
)

const (
	defaultOrder = orderedbytes.Ascending
)

func toMillis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func marshal(value any) ([]byte, error) {
	var b []byte
	switch v := value.(type) {
	case time.Time:
		b = orderedbytes.EncodeInt64(toMillis(v), defaultOrder)
	case bool:
		val := int8(0)
		if v {
			val = math.MaxInt8
		}
		b = orderedbytes.EncodeInt8(val, defaultOrder)
	default:
		var err error
		b, err = orderedbytes.Encode(v, defaultOrder)
		if err != nil {
			return nil, err
		}
	}
	return b, nil
}

func mustMarshal(value any) []byte {
	b, err := marshal(value)
	if err != nil {
		panic(err)
	}
	return b
}
