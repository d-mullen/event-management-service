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

package protobufutils

import (
	"sort"

	"github.com/pkg/errors"
	"github.com/spaolacci/murmur3"
	"github.com/zenoss/zing-proto/v11/go/cloud/common"
	"github.com/zenoss/zingo/v4/orderedbytes"
)

func MustFromScalarArray(a *common.ScalarArray) []interface{} {
	result, err := FromScalarArray(a)
	if err != nil {
		panic(errors.WithStack(err).Error())
	}
	return result
}

func FromScalarArray(a *common.ScalarArray) ([]interface{}, error) {
	if a == nil {
		return nil, nil
	}
	arrayValues := make([]interface{}, len(a.Scalars))
	for i, av := range a.Scalars {
		value, err := FromScalar(av)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		arrayValues[i] = value
	}
	return arrayValues, nil
}

func MustToScalarArray(a []interface{}) *common.ScalarArray {
	result, err := ToScalarArray(a)
	if err != nil {
		panic(errors.WithStack(err).Error())
	}
	return result
}

func ToScalarArray(a []interface{}) (*common.ScalarArray, error) {
	scalarArray := make([]*common.Scalar, len(a))
	for i, v := range a {
		scalarV, err := ToScalar(v)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert interface to Scalar")
		}
		scalarArray[i] = scalarV
	}
	return &common.ScalarArray{Scalars: scalarArray}, nil
}

func MustFromScalarArrayMap(aaMap map[string]*common.ScalarArray) map[string][]interface{} {
	result, err := FromScalarArrayMap(aaMap)
	if err != nil {
		panic(errors.WithStack(err).Error())
	}
	return result
}

func FromScalarArrayMap(aaMap map[string]*common.ScalarArray) (map[string][]interface{}, error) {
	result := make(map[string][]interface{}, len(aaMap))
	for k, av := range aaMap {
		v, err := FromScalarArray(av)
		if err != nil {
			return nil, errors.Wrapf(err, "error converting array with key %s", k)
		}
		result[k] = v
	}
	return result, nil
}

func GetStringListFromScalarArrayMap(aaMap map[string]*common.ScalarArray, key string) ([]string, error) {
	scArray, ok := aaMap[key]
	if !ok {
		return nil, nil
	}

	rawValues, err := FromScalarArray(scArray)
	if err != nil {
		return nil, err
	}

	values := make([]string, len(rawValues))
	for i, rawValue := range rawValues {
		value, ok := rawValue.(string)
		if ok {
			values[i] = value
		}
	}

	return values, nil
}

func GetStringValueFromScalarArrayMap(aaMap map[string]*common.ScalarArray, key string) (string, error) {
	values, err := GetStringListFromScalarArrayMap(aaMap, key)
	if err != nil {
		return "", err
	}

	if len(values) > 0 {
		return values[0], nil
	}
	return "", nil
}

func MustToScalarArrayMap(imap map[string][]interface{}) map[string]*common.ScalarArray {
	result, err := ToScalarArrayMap(imap)
	if err != nil {
		panic(errors.WithStack(err).Error())
	}
	return result
}

func ToScalarArrayMap(imap map[string][]interface{}) (map[string]*common.ScalarArray, error) {
	aamap := make(map[string]*common.ScalarArray, len(imap))
	for k, v := range imap {
		aarray, err := ToScalarArray(v)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to convert array with key %s to ScalarArray", k)
		}
		aamap[k] = aarray
	}
	return aamap, nil
}

// OrderedScalarSlicesAreEqual checks to see if two scalar slices are equal.
func OrderedScalarSlicesAreEqual(array1 []*common.Scalar, array2 []*common.Scalar) bool {
	if array1 == nil {
		return array2 == nil
	}
	if array2 == nil || len(array1) != len(array2) {
		return false
	}

	for index, value := range array1 {
		if !ScalarsAreEqual(value, array2[index]) {
			return false
		}
	}
	return true
}

// ScalarSlicesAreEqual checks to see if two scalar slices are equal, regardless of elements order.
func ScalarSlicesAreEqual(array1 []*common.Scalar, array2 []*common.Scalar) bool {
	if array1 == nil {
		return array2 == nil
	}
	if array2 == nil || len(array1) != len(array2) {
		return false
	}

	a1Copy := ScalarArraySort(array1)
	a2Copy := ScalarArraySort(array2)

	return OrderedScalarSlicesAreEqual(a1Copy, a2Copy)
}

// FindScalarInScalarArray checks to see if a scalar array contains a given scalar value.
func FindScalarInScalarArray(scalar *common.Scalar, array []*common.Scalar) bool {
	for _, value := range array {
		if ScalarsAreEqual(value, scalar) {
			return true
		}
	}
	return false
}

// ScalarArraySort returns a sorted copy of the scalar array.
func ScalarArraySort(arr []*common.Scalar) []*common.Scalar {
	if arr == nil || len(arr) == 0 {
		return arr
	}

	arrCopy := make([]*common.Scalar, len(arr))
	copy(arrCopy, arr)
	sort.Slice(arrCopy, func(i, j int) bool {
		return arrCopy[i].String() < arrCopy[j].String()
	})

	return arrCopy
}

// ScalarArrayMapsAreEqual checks to see if two scalar array maps are equal,
// regardless of elements order.
func ScalarArrayMapsAreEqual(a map[string]*common.ScalarArray, b map[string]*common.ScalarArray) bool {
	if a == nil {
		return b == nil
	}
	if b == nil || len(a) != len(b) {
		return false
	}

	for keyAi, arrA := range a {
		if arrB, ok := b[keyAi]; ok {
			if !ScalarSlicesAreEqual(arrA.Scalars, arrB.Scalars) {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// ScalarArrayMapHash computes a hash of scalar array map, resilient to sort order.
func ScalarArrayMapHash(scalarMap map[string]*common.ScalarArray) uint64 {
	hasher := murmur3.New64()

	fields := make([]string, 0, len(scalarMap))
	for field := range scalarMap {
		fields = append(fields, field)
	}
	sort.Strings(fields)

	for _, f := range fields {
		hasher.Write([]byte(f))

		scalarArrCopy := ScalarArraySort(scalarMap[f].Scalars)
		for _, val := range scalarArrCopy {
			value, _ := FromScalar(val)
			b, _ := orderedbytes.Encode(value, orderedbytes.Ascending)
			hasher.Write(b)
		}
	}

	return hasher.Sum64()
}
