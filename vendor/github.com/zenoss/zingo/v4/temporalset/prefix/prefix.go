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
package prefix

import (
	"math"
	"math/bits"
)

// Table of masks, to save operations later
var (
	prefixMasks [65]uint64
	bitMasks    [65]uint64
)

func init() {
	// Populate the mask tables
	for i := uint(0); i < 64; i++ {
		prefixMasks[i] = math.MaxUint64 << i
	}
	for i := uint(1); i < 64; i++ {
		bitMasks[i] = 1 << (i - 1)
	}
}

// BranchingBit returns the level of the bit (higher == more
// significant) at which this prefix and the provided prefix begin to
// differ.
func BranchingBit(a, b uint64) uint {
	return uint(bits.Len64(a ^ b))
}

// MaskAbove returns the prefix above a given bit.
func MaskAbove(a uint64, level uint) uint64 {
	return a & prefixMasks[level]
}

// ZeroAt returns whether the bit at the provided level is zero.
func ZeroAt(a uint64, level uint) bool {
	return a&bitMasks[level] == 0
}

// IsPrefixAt returns whether this prefix is a prefix of the given prefix at
// the given level.
func IsPrefixAt(a, prefix uint64, level uint) bool {
	return MaskAbove(a, level) == prefix
}
