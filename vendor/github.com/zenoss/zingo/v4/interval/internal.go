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

import "math"

func toClosedLower(b uint64, bt BoundType) (uint64, int) {
	switch bt {
	case Unbound:
		return 0, 0
	case ClosedBound:
		return b, 2
	case OpenBound:
		return b + 1, 1
	case EmptyBound:
		panic("can't compare empty bound")
	default:
		panic("unknown bound type")
	}
}

func toClosedUpper(b uint64, bt BoundType) (uint64, int) {
	switch bt {
	case Unbound:
		return math.MaxUint64, 2
	case ClosedBound:
		return b, 0
	case OpenBound:
		return b - 1, 1
	case EmptyBound:
		panic("can't compare empty bound")
	default:
		panic("unknown bound type")
	}
}

// Determines whether lhs's lower bound is lower than rhs's lower bound.
//  Does not handle empty intervals, you must check for IsEmpty first
func hasLowerLowerBound(lhb uint64, lhbt BoundType, rhb uint64, rhbt BoundType) bool {
	// Handle edge cases
	if lhb == math.MaxUint64 && lhbt == OpenBound {
		return false
	}
	if rhb == math.MaxUint64 && rhbt == OpenBound {
		return true
	}

	lhsb, lhsbt := toClosedLower(lhb, lhbt)
	rhsb, rhsbt := toClosedLower(rhb, rhbt)

	if lhsb == rhsb {
		return lhsbt < rhsbt
	}

	return lhsb < rhsb

}

// Determines whether lhs's upper bound is lower than rhs's upper bound.
//  Does not handle empty intervals, you must check for IsEmpty first
func hasLowerUpperBound(lhb uint64, lhbt BoundType, rhb uint64, rhbt BoundType) bool {

	// Handle edge cases
	if rhb == 0 && rhbt == OpenBound {
		return false
	}
	if lhb == 0 && lhbt == OpenBound {
		return true
	}

	lhsb, lhsbt := toClosedUpper(lhb, lhbt)
	rhsb, rhsbt := toClosedUpper(rhb, rhbt)

	if lhsb == rhsb {
		return lhsbt < rhsbt
	}

	return lhsb < rhsb
}

// isEmpty returns whether the bound values passed represent an empty interval
func isEmpty(lbt BoundType, lb, ub uint64, ubt BoundType) bool {
	if lbt == EmptyBound || ubt == EmptyBound {
		return true
	}
	if lbt == Unbound || ubt == Unbound {
		return false
	}
	if lb > ub {
		return true
	}
	return lb == ub && (lbt == OpenBound || ubt == OpenBound)
}

func isUnbounded(lbt, ubt BoundType) bool {
	return lbt == Unbound && ubt == Unbound
}

func isHalfUnbounded(lbt, ubt BoundType) bool {
	return lbt == Unbound || ubt == Unbound
}

func boundsEqual(lb uint64, lbt BoundType, rb uint64, rbt BoundType) bool {
	return lbt == rbt && (lbt == Unbound || lb == rb)
}

func intersection(llb uint64, llbt BoundType, lub uint64, lubt BoundType, rlb uint64, rlbt BoundType, rub uint64, rubt BoundType) (lbt BoundType, lb uint64, ub uint64, ubt BoundType) {
	if isEmpty(llbt, llb, lub, lubt) || isEmpty(rlbt, rlb, rub, rubt) {
		return EmptyBound, 0, 0, EmptyBound
	}
	if hasLowerLowerBound(llb, llbt, rlb, rlbt) {
		lb, lbt = rlb, rlbt
	} else {
		lb, lbt = llb, llbt
	}
	if hasLowerUpperBound(lub, lubt, rub, rubt) {
		ub, ubt = lub, lubt
	} else {
		ub, ubt = rub, rubt
	}
	return
}

func encompass(llb uint64, llbt BoundType, lub uint64, lubt BoundType, rlb uint64, rlbt BoundType, rub uint64, rubt BoundType) (lbt BoundType, lb uint64, ub uint64, ubt BoundType) {
	if isEmpty(rlbt, rlb, rub, rubt) {
		return llbt, llb, lub, lubt
	}
	if isEmpty(llbt, llb, lub, lubt) {
		return rlbt, rlb, rub, rubt
	}
	if hasLowerLowerBound(llb, llbt, rlb, rlbt) {
		lb, lbt = llb, llbt
	} else {
		lb, lbt = rlb, rlbt
	}
	if hasLowerUpperBound(lub, lubt, rub, rubt) {
		ub, ubt = rub, rubt
	} else {
		ub, ubt = lub, lubt
	}
	return
}

func contains(llb uint64, llbt BoundType, lub uint64, lubt BoundType, rlb uint64, rlbt BoundType, rub uint64, rubt BoundType) bool {
	if isUnbounded(llbt, lubt) || isEmpty(rlbt, rlb, rub, rubt) {
		return true
	}
	if isUnbounded(rlbt, rubt) || isEmpty(llbt, llb, lub, lubt) {
		return false
	}
	lower := !hasLowerLowerBound(rlb, rlbt, llb, llbt)
	upper := !hasLowerUpperBound(lub, lubt, rub, rubt)
	return lower && upper
}
