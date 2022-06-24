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
	"github.com/zenoss/zingo/v4/interval"
)

type (
	trie struct {
		prefix        uint64
		level         uint
		unbounded     bool
		isLeaf        bool
		includesValue bool
		left, right   int
		nextUnused    int
	}
)

func (t trie) Prefix() uint64 {
	return t.prefix
}

func (t trie) Level() uint {
	return t.level
}

func (t trie) Unbounded() bool {
	return t.unbounded
}

func (t trie) branch(lset, rset *IntervalSet, branch, left, right int) int {
	if branch != left && branch != right {
		// we are replacing branch, de-allocate this slot only
		lset.deAllocate(branch)
	}
	result := maybeNewBranch(lset, rset, t.Prefix(), t.Level(), left, right)
	return result
}

func lowerBound(s *IntervalSet, idx int) (uint64, interval.BoundType) {
	node := s.nodes[idx]
	if node.isLeaf {
		return node.leafLower()
	}
	return lowerBound(s, node.left)
}

func upperBound(s *IntervalSet, idx int) (uint64, interval.BoundType) {
	node := s.nodes[idx]
	if node.isLeaf {
		return node.leafUpper()
	}
	return upperBound(s, node.right)
}

func (t trie) leafPreceding(s *IntervalSet, key uint64, ul bool) (trie, bool, bool) {
	if t.isLeaf {
		return t.leafLeafPreceding(key, ul)
	}
	return t.branchLeafPreceding(s, key, ul)
}

func (t trie) leafSucceeding(s *IntervalSet, key uint64, ul bool) (trie, bool, bool) {
	if t.isLeaf {
		return t.leafLeafSucceeding(key, ul)
	}
	return t.branchLeafSucceeding(s, key, ul)
}

func (t trie) rightmostLeaf(s *IntervalSet, ul bool) (trie, bool) {
	if t.isLeaf {
		return t, ul
	} else {
		return s.nodes[t.right].rightmostLeaf(s, ul != s.nodes[t.left].Unbounded())
	}
}

func (t trie) leftmostLeaf(s *IntervalSet, ul bool) (trie, bool) {
	if t.isLeaf {
		return t, ul
	} else {
		return s.nodes[t.left].leftmostLeaf(s, ul)
	}
}

func (t trie) iterate(s *IntervalSet, f func(key uint64, left bool, bound interval.BoundType), prevkey uint64, prevleft bool, prevbound interval.BoundType, unboundedLeft bool) (uint64, bool, interval.BoundType) {
	if t.isLeaf {
		return t.leafIterate(f, prevkey, prevleft, prevbound, unboundedLeft)
	} else {
		return t.branchIterate(s, f, prevkey, prevleft, prevbound, unboundedLeft)
	}
}

func trieEq(as, bs *IntervalSet, aidx, bidx int) bool {
	an, bn := aidx < 0, bidx < 0
	if an != bn {
		return false
	}
	if an {
		return true
	}
	a, b := as.nodes[aidx], bs.nodes[bidx]
	switch a.isLeaf {
	case true:
		return a.prefix == b.prefix &&
			a.includesValue == b.includesValue &&
			a.isLeaf == b.isLeaf &&
			a.unbounded == b.unbounded
	case false:
		if b.isLeaf {
			return false
		}
		if a.level != b.level {
			return false
		}
		if a.prefix != b.prefix {
			return false
		}
		if !trieEq(as, bs, a.left, b.left) || !trieEq(as, bs, a.right, b.right) {
			return false
		}
	}
	return true
}
