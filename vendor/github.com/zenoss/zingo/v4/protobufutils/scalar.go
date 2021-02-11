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
	"github.com/zenoss/zing-proto/v11/go/cloud/common"
)

var (
	ErrInvalidScalarType = errors.New("invalid scalar type")
)

func MustToScalar(d interface{}) *common.Scalar {
	ne, err := ToScalar(d)
	if err != nil {
		panic(errors.WithStack(err).Error())
	}
	return ne
}

func ToScalar(d interface{}) (*common.Scalar, error) {
	switch d := d.(type) {
	case bool:
		return &common.Scalar{
			Value: &common.Scalar_BoolVal{
				d,
			},
		}, nil
	case int:
		cd := int64(d)
		return &common.Scalar{
			Value: &common.Scalar_LongVal{
				cd,
			},
		}, nil
	case int8:
		cd := int64(d)
		return &common.Scalar{
			Value: &common.Scalar_LongVal{
				cd,
			},
		}, nil
	case int16:
		cd := int64(d)
		return &common.Scalar{
			Value: &common.Scalar_LongVal{
				cd,
			},
		}, nil
	case int32:
		cd := int64(d)
		return &common.Scalar{
			Value: &common.Scalar_LongVal{
				cd,
			},
		}, nil
	case int64:
		return &common.Scalar{
			Value: &common.Scalar_LongVal{
				d,
			},
		}, nil
	case uint:
		cd := uint32(d)
		return &common.Scalar{
			Value: &common.Scalar_UintVal{
				cd,
			},
		}, nil
	case uint8:
		cd := uint32(d)
		return &common.Scalar{
			Value: &common.Scalar_UintVal{
				cd,
			},
		}, nil
	case uint16:
		cd := uint32(d)
		return &common.Scalar{
			Value: &common.Scalar_UintVal{
				cd,
			},
		}, nil
	case uint32:
		return &common.Scalar{
			Value: &common.Scalar_UintVal{
				d,
			},
		}, nil
	case uint64:
		return &common.Scalar{
			Value: &common.Scalar_UlongVal{
				d,
			},
		}, nil
	case float32:
		return &common.Scalar{
			Value: &common.Scalar_FloatVal{
				d,
			},
		}, nil
	case float64:
		return &common.Scalar{
			Value: &common.Scalar_DoubleVal{
				d,
			},
		}, nil
	case string:
		return &common.Scalar{
			Value: &common.Scalar_StringVal{
				d,
			},
		}, nil
	default:
		return nil, errors.Wrapf(ErrInvalidScalarType, "could not convert %v to Scalar", d)
	}
}

func MustFromScalar(a *common.Scalar) interface{} {
	ne, err := FromScalar(a)
	if err != nil {
		panic(errors.WithStack(err).Error())
	}
	return ne
}

func FromScalar(a *common.Scalar) (interface{}, error) {
	switch d := a.GetValue().(type) {
	case *common.Scalar_LongVal:
		return d.LongVal, nil
	case *common.Scalar_UlongVal:
		return d.UlongVal, nil
	case *common.Scalar_UintVal:
		return d.UintVal, nil
	case *common.Scalar_FloatVal:
		return d.FloatVal, nil
	case *common.Scalar_DoubleVal:
		return d.DoubleVal, nil
	case *common.Scalar_StringVal:
		return d.StringVal, nil
	case *common.Scalar_BoolVal:
		return d.BoolVal, nil
	}
	return nil, errors.Wrapf(ErrInvalidScalarType, "error converting %v to interface", a)
}

func MustToScalarMap(m map[string]interface{}) map[string]*common.Scalar {
	result, err := ToScalarMap(m)
	if err != nil {
		panic(errors.WithStack(err).Error())
	}
	return result
}

func ToScalarMap(m map[string]interface{}) (map[string]*common.Scalar, error) {
	result := make(map[string]*common.Scalar, len(m))
	for k, av := range m {
		v, err := ToScalar(av)
		if err == nil {
			result[k] = v
		} else {
			return nil, errors.Wrapf(err, "Could not encode data[%s] to protobuf", k)
		}
	}
	return result, nil
}

func MustFromScalarMap(scalarMap map[string]*common.Scalar) map[string]interface{} {
	result, err := FromScalarMap(scalarMap)
	if err != nil {
		panic(errors.WithStack(err).Error())
	}
	return result
}

func FromScalarMap(scalarMap map[string]*common.Scalar) (map[string]interface{}, error) {
	result := make(map[string]interface{}, len(scalarMap))
	for k, av := range scalarMap {
		v, err := FromScalar(av)
		if err != nil {
			return nil, errors.Wrapf(err, "Could not decode data[%s] from Scalar", k)
		}
		result[k] = v
	}
	return result, nil
}

// This should be the fastest way to tell if two Scalars are equal
func ScalarsAreEqual(scalar1, scalar2 *common.Scalar) bool {
	// if they are both nil, they are equal
	if scalar1 == nil && scalar2 == nil {
		return true
	}

	// if only one is nil, they are not equal
	if scalar1 == nil || scalar2 == nil {
		return false
	}

	// if they point to the same object, they are equal
	if scalar1 == scalar2 {
		return true
	}

	return scalar1.GetBoolVal() == scalar2.GetBoolVal() &&
		scalar1.GetLongVal() == scalar2.GetLongVal() &&
		scalar1.GetUlongVal() == scalar2.GetUlongVal() &&
		scalar1.GetUintVal() == scalar2.GetUintVal() &&
		scalar1.GetFloatVal() == scalar2.GetFloatVal() &&
		scalar1.GetDoubleVal() == scalar2.GetDoubleVal() &&
		scalar1.GetStringVal() == scalar2.GetStringVal()
}
