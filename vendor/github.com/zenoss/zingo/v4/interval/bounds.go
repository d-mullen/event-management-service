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
package interval

func openBounds(lower, upper uint64) (BoundType, uint64, uint64, BoundType) {
	return OpenBound, lower, upper, OpenBound
}

func closedBounds(lower, upper uint64) (BoundType, uint64, uint64, BoundType) {
	return ClosedBound, lower, upper, ClosedBound
}

func leftOpenBounds(lower, upper uint64) (BoundType, uint64, uint64, BoundType) {
	return OpenBound, lower, upper, ClosedBound
}

func rightOpenBounds(lower, upper uint64) (BoundType, uint64, uint64, BoundType) {
	return ClosedBound, lower, upper, OpenBound
}

func emptyBounds() (BoundType, uint64, uint64, BoundType) {
	return EmptyBound, 0, 0, EmptyBound
}

func unboundedBounds() (BoundType, uint64, uint64, BoundType) {
	return Unbound, 0, 0, Unbound
}

func belowBounds(upper uint64) (BoundType, uint64, uint64, BoundType) {
	return Unbound, 0, upper, OpenBound
}

func atOrBelowBounds(upper uint64) (BoundType, uint64, uint64, BoundType) {
	return Unbound, 0, upper, ClosedBound
}

func aboveBounds(lower uint64) (BoundType, uint64, uint64, BoundType) {
	return OpenBound, lower, 0, Unbound
}

func atOrAboveBounds(lower uint64) (BoundType, uint64, uint64, BoundType) {
	return ClosedBound, lower, 0, Unbound
}

func pointBounds(i uint64) (BoundType, uint64, uint64, BoundType) {
	return ClosedBound, i, i, ClosedBound
}
