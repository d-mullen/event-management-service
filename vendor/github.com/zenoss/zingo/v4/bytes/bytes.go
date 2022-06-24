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
package bytes

import (
	"encoding/binary"
	"math"
	"reflect"
	"runtime"
	"unsafe"
)

// Byte conversions that match the java hbase Bytes package

func IntToBytes(v int) []byte {
	return Int32ToBytes(int32(v))
}

func BytesToInt(b []byte) int {
	return int(BytesToInt32(b))
}

func Int32ToBytes(v int32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(v))
	return b
}

func BytesToInt32(b []byte) int32 {
	v := binary.BigEndian.Uint32(b)
	return int32(v)
}

func Int64ToBytes(v int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func BytesToInt64(b []byte) int64 {
	v := binary.BigEndian.Uint64(b)
	return int64(v)
}

func BoolToBytes(v bool) []byte {
	var b byte
	if v {
		b = byte(255)
	} else {
		b = byte(0)
	}
	return []byte{b}
}

func BytesToBool(b []byte) bool {
	return b[0] != byte(0)
}

func StringToBytes(v string) []byte {
	return []byte(v)
}

func BytesToString(b []byte) string {
	return string(b)
}

// StringToBytesUnsafe produces a read-only byte slice pointing to the data in
// memory backing the given string (i.e, without a copy). Modifying the
// returned slice will panic.
func StringToBytesUnsafe(s string) (b []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	runtime.KeepAlive(&s)
	return
}

func BytesToFloat32(data []byte) float32 {
	bits := binary.BigEndian.Uint32(data)
	return math.Float32frombits(bits)
}

func Float32ToBytes(f float32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, math.Float32bits(f))
	return b
}

func BytesToFloat64(data []byte) float64 {
	bits := binary.BigEndian.Uint64(data)
	return math.Float64frombits(bits)
}

func Float64ToBytes(f float64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, math.Float64bits(f))
	return b
}
