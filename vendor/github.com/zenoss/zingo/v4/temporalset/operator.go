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
package temporalset

import (
	"github.com/zenoss/zingo/v4/temporalset/prefix"
)

var (
	treeOr  = operator(orCollision, orOverlapLhs, orOverlapRhs)
	treeAnd = operator(andCollision, andOverlapLhs, andOverlapRhs)
	treeXor = operator(xorCollision, xorOverlapLhs, xorOverlapRhs)
)

type (
	collisionFunc func(*IntervalSet, *IntervalSet, bool, bool, int, int) int
	overlapFunc   func(*IntervalSet, *IntervalSet, bool, bool, int) int
)

// Operator returns a function that performs an operation on two temporalsets (lset and rset) on the nodes at the specified
//  indices (lhs, rhs).  The result is always captured in lset, and the return value is the index of a node in lset.
func operator(c collisionFunc, ol overlapFunc, or overlapFunc) func(*IntervalSet, *IntervalSet, bool, bool, int, int) int {
	return func(lset, rset *IntervalSet, lul, rul bool, lhs, rhs int) int {
		if lhs < 0 && rhs < 0 {
			return -1
		}
		if lhs < 0 {
			orResult := or(lset, rset, lul, rul, rhs)
			return orResult
		}
		if rhs < 0 {
			olResult := ol(lset, lset, lul, rul, lhs)
			return olResult
		}
		result := performOperation(lset, rset, c, ol, or, lul, rul, lhs, rhs)
		return result
	}
}

func performOperation(lset, rset *IntervalSet, collision collisionFunc, overlapLhs overlapFunc, overlapRhs overlapFunc, lul, rul bool, lhsidx, rhsidx int) int {
	lhs := lset.nodes[lhsidx]
	rhs := rset.nodes[rhsidx]

	ap, bp := lhs.Prefix(), rhs.Prefix()
	al, bl := lhs.Level(), rhs.Level()
	if al > bl {
		if !prefix.IsPrefixAt(bp, ap, al) {
			// a and b don't share a prefix, so just join the trees
			return join(lset, rset, overlapLhs, overlapRhs, lul, rul, lhsidx, rhsidx)
		}
		// lhs has a larger level, so it has to be a branch
		aul := lul != lset.nodes[lhs.left].Unbounded()
		var lh1, rh1 int
		if prefix.ZeroAt(bp, al) {
			// rhs is somewhere within the left child of lhs; recurse
			lh1 = performOperation(lset, rset, collision, overlapLhs, overlapRhs, lul, rul, lhs.left, rhsidx)
			rh1 = overlapLhs(lset, lset, aul, rul != rhs.Unbounded(), lhs.right)
		} else {
			// rhs is somewhere within the right child of lhs; recurse
			lh1 = overlapLhs(lset, lset, lul, rul, lhs.left)
			rh1 = performOperation(lset, rset, collision, overlapLhs, overlapRhs, aul, rul, lhs.right, rhsidx)
		}
		return lhs.branch(lset, lset, lhsidx, lh1, rh1)
	} else if bl > al {
		if !prefix.IsPrefixAt(ap, bp, bl) {
			// a and b don't share a prefix, so just join the trees
			return join(lset, rset, overlapLhs, overlapRhs, lul, rul, lhsidx, rhsidx)
		}
		// rhs has a larger level, so it has to be a branch
		bul := rul != rset.nodes[rhs.left].Unbounded()
		var lh1, rh1 int
		if prefix.ZeroAt(ap, bl) {
			// lhs is somewhere within the left child of rhs; recurse
			lh1 = performOperation(lset, rset, collision, overlapLhs, overlapRhs, lul, rul, lhsidx, rhs.left)
			rh1 = overlapRhs(lset, rset, lul != lhs.Unbounded(), bul, rhs.right)
		} else {
			// rhs is somewhere within the right child of lhs; recurse
			lh1 = overlapRhs(lset, rset, lul, rul, rhs.left)
			rh1 = performOperation(lset, rset, collision, overlapLhs, overlapRhs, lul, bul, lhsidx, rhs.right)
		}
		// At this point, both lh1 and rh1 are indices in lset
		return rhs.branch(lset, lset, -1, lh1, rh1)
	} else {
		// The trees are at the same level
		if ap != bp {
			// Same level, different prefix
			r := join(lset, rset, overlapLhs, overlapRhs, lul, rul, lhsidx, rhsidx)
			return r
		}
		// Because they have the same level, they are either both branches or
		// both leaves
		if lhs.isLeaf {
			return collision(lset, rset, lul, rul, lhsidx, rhsidx)
		} else {
			aul := lul != lset.nodes[lhs.left].Unbounded()
			bul := rul != rset.nodes[rhs.left].Unbounded()
			lhs1 := performOperation(lset, rset, collision, overlapLhs, overlapRhs, lul, rul, lhs.left, rhs.left)
			rhs1 := performOperation(lset, rset, collision, overlapLhs, overlapRhs, aul, bul, lhs.right, rhs.right)
			return lhs.branch(lset, lset, lhsidx, lhs1, rhs1)
		}
	}
}

// join creates a new tree that encompasses the given trees
//  It will modify lset as necessary, and return an index from lset that points to the root of the resulting tree
func join(lset, rset *IntervalSet, overlapLhs overlapFunc, overlapRhs overlapFunc, lul, rul bool, lhsidx, rhsidx int) int {

	lhs, rhs := lset.nodes[lhsidx], rset.nodes[rhsidx]
	ap, bp := lhs.Prefix(), rhs.Prefix()
	level := prefix.BranchingBit(ap, bp)
	p := prefix.MaskAbove(ap, level)

	var lhs1, rhs1 int
	if prefix.ZeroAt(ap, level) {
		lhs1 = overlapLhs(lset, lset, lul, rul, lhsidx)
		rhs1 = overlapRhs(lset, rset, lul != lhs.Unbounded(), rul, rhsidx)
	} else {
		lhs1 = overlapRhs(lset, rset, lul, rul, rhsidx)
		rhs1 = overlapLhs(lset, lset, lul, rul != rhs.Unbounded(), lhsidx)
	}
	result := maybeNewBranch(lset, lset, p, level, lhs1, rhs1)

	return result
}

func orCollision(lset, rset *IntervalSet, lul, rul bool, lhsidx, rhsidx int) int {
	lhs, rhs := lset.nodes[lhsidx], rset.nodes[rhsidx]
	below := lul || rul
	includes := (lhs.includesValue != lul) || (rhs.includesValue != rul)
	above := (lhs.unbounded != lul) || (rhs.unbounded != rul)
	result := maybeNewLeaf(lset, rset, below != includes, includes != above, lhsidx, rhsidx)
	return result
}

func orOverlapLhs(dst, src *IntervalSet, lul, rul bool, lhs int) int {
	if rul {
		if dst == src {
			// De-allocate lhs
			dst.removeNode(lhs)
		}
		return -1
	}
	if dst != src {
		// copy the node from src to dst
		lhs = dst.copyNodes(src, lhs)
	}
	return lhs
}

func orOverlapRhs(dst, src *IntervalSet, lul, rul bool, rhs int) int {
	if lul {
		if dst == src {
			// De-allocate rhs
			dst.removeNode(rhs)
		}
		return -1
	}
	if dst != src {
		// copy the node from src to dst
		rhs = dst.copyNodes(src, rhs)
	}
	return rhs
}

func andCollision(lset, rset *IntervalSet, lul, rul bool, lhsidx, rhsidx int) int {
	lhs, rhs := lset.nodes[lhsidx], rset.nodes[rhsidx]
	below := lul && rul
	includes := (lhs.includesValue != lul) && (rhs.includesValue != rul)
	above := (lhs.unbounded != lul) && (rhs.unbounded != rul)
	result := maybeNewLeaf(lset, rset, below != includes, includes != above, lhsidx, rhsidx)
	return result
}

func andOverlapLhs(dst, src *IntervalSet, lul, rul bool, lhs int) int {
	if rul {
		if dst != src {
			// copy the node from src to dst
			lhs = dst.copyNodes(src, lhs)
		}
		return lhs
	}
	if dst == src {
		// De-allocate lhs
		dst.removeNode(lhs)
	}
	return -1
}

func andOverlapRhs(dst, src *IntervalSet, lul, rul bool, rhs int) int {
	if lul {
		if dst != src {
			// copy the node from src to dst
			rhs = dst.copyNodes(src, rhs)
		}
		return rhs
	}
	if dst == src {
		// De-allocate rhs
		dst.removeNode(rhs)
	}
	return -1
}

func xorCollision(lset, rset *IntervalSet, lul, rul bool, lhsidx, rhsidx int) int {
	lhs, rhs := lset.nodes[lhsidx], rset.nodes[rhsidx]
	below := lul != rul
	includes := (lhs.includesValue != lul) != (rhs.includesValue != rul)
	above := (lhs.unbounded != lul) != (rhs.unbounded != rul)
	result := maybeNewLeaf(lset, rset, below != includes, includes != above, lhsidx, rhsidx)
	return result
}

func xorOverlapLhs(dst, src *IntervalSet, lul, rul bool, lhs int) int {
	if dst != src {
		// copy the node from src to dst
		lhs = dst.copyNodes(src, lhs)
	}
	return lhs
}

func xorOverlapRhs(dst, src *IntervalSet, lul, rul bool, rhs int) int {
	if dst != src {
		// copy the node from src to dst
		rhs = dst.copyNodes(src, rhs)
	}
	return rhs
}
