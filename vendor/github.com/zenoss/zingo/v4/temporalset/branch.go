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
	"github.com/zenoss/zingo/v4/temporalset/prefix"
)

func (t trie) branchLeafPreceding(s *IntervalSet, key uint64, ul bool) (trie, bool, bool) {
	left := s.nodes[t.left]
	right := s.nodes[t.right]
	if prefix.IsPrefixAt(key, t.Prefix(), t.Level()) {
		// This node is in the path! Keep traversing.
		if prefix.ZeroAt(key, t.level) {
			return left.leafPreceding(s, key, ul)
		}
		right, rul, ok := right.leafPreceding(s, key, ul != left.Unbounded())
		if !ok {
			l, ul := left.rightmostLeaf(s, ul)
			return l, ul, true
		}
		return right, rul, true
	}
	// See which side of our key this node falls on
	if t.Prefix() > key {
		// This trie doesn't have any leaves to the left of our key. Turn
		// around and take the next left.
		return trie{}, false, false
	}
	// The key is to the right of this prefix, so we want the rightmost leaf of
	// this trie.
	l, ul := t.rightmostLeaf(s, ul)
	return l, ul, true
}

func (t trie) branchLeafSucceeding(s *IntervalSet, key uint64, ul bool) (trie, bool, bool) {
	tleft, tright := s.nodes[t.left], s.nodes[t.right]
	if prefix.IsPrefixAt(key, t.Prefix(), t.Level()) {
		// This node is in the path! Keep traversing.
		if prefix.ZeroAt(key, t.level) {
			left, lul, ok := tleft.leafSucceeding(s, key, ul)
			if !ok {
				l, ul := tright.leftmostLeaf(s, ul != tleft.Unbounded())
				return l, ul, true
			}
			return left, lul, true
		}
		return tright.leafSucceeding(s, key, ul != tleft.Unbounded())
	}
	// See which side of our key this node falls on
	if t.Prefix() < key {
		// This trie doesn't have any leaves to the right of our key. All done
		// here
		return trie{}, false, false
	}
	// The key is to the left of this prefix, so we want the leftmost leaf of
	// this trie.
	l, ul := t.leftmostLeaf(s, ul)
	return l, ul, true
}

func (t trie) branchIterate(s *IntervalSet, f func(key uint64, left bool, bound interval.BoundType), prevkey uint64, prevleft bool, prevbound interval.BoundType, unboundedLeft bool) (uint64, bool, interval.BoundType) {
	left, right := s.nodes[t.left], s.nodes[t.right]
	am := unboundedLeft != left.Unbounded()
	pk, pl, pb := left.iterate(s, f, prevkey, prevleft, prevbound, unboundedLeft)
	return right.iterate(s, f, pk, pl, pb, am)
}

// Creates a new branch, if necessary, in lset and returns the index in lset of that branch
func maybeNewBranch(lset, rset *IntervalSet, p uint64, level uint, left, right int) int {
	if left < 0 && right < 0 {
		return -1
	}
	if left < 0 {
		if lset == rset {
			return right
		}
		// Copy the sub-tree rooted at right from rset to lset
		return lset.copyNodes(rset, right)
	}
	if right < 0 {
		return left
	}

	if lset != rset {
		// Copy the sub-tree rooted at right from rset to lset
		right = lset.copyNodes(rset, right)
	}

	newBranch := trie{
		prefix:     p,
		level:      level,
		unbounded:  lset.nodes[left].Unbounded() != rset.nodes[right].Unbounded(),
		left:       left,
		right:      right,
		nextUnused: -1,
	}
	return lset.addNode(newBranch)
}
