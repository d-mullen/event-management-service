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

func RightOpenSplitN(ival Interval, n int) []Interval {
	return splitN(ival, n, false)
}

func LeftOpenSplitN(ival Interval, n int) []Interval {
	return splitN(ival, n, true)

}

func splitN(ival Interval, n int, left bool) (ivals []Interval) {
	switch {
	case n == 0, ival.IsEmpty(), ival.IsHalfUnbounded():
		return
	case n == 1:
		ivals = append(ivals, ival)
		return
	}
	span := float64(ival.Span()) / float64(n)
	offset, _ := ival.Lower()
	ivals = make([]Interval, n)
	var subtrahend Interval
	for i := 0; i < n-1; i++ {
		next := offset + uint64(span*float64(i+1))
		switch {
		case left:
			subtrahend = Above(next)
		default:
			subtrahend = AtOrAbove(next)
		}
		ivals[i], ival, _ = ival.Partition(subtrahend)
	}
	ivals[n-1] = ival
	return
}
