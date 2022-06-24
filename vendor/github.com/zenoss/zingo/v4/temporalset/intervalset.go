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
	"strings"

	"github.com/zenoss/zingo/v4/interval"
)

type IntervalSet struct {
	root      int
	unused    int
	nodes     []trie
	ul        bool
	nodeCount int
}

// init makes sure s.root is initialized to -1 when there are no nodes
//  This is necessary to ensure IntervalSets created without a constructor are treated as empty
func (s *IntervalSet) init() {
	if s != nil && len(s.nodes) == 0 {
		s.root = -1
		s.unused = -1
	}
}

func (s *IntervalSet) Size() int {
	return s.nodeCount
}

func (s *IntervalSet) addNode(t trie) int {
	s.nodeCount++
	if s.unused >= 0 {
		idx := s.unused
		s.unused = s.nodes[idx].nextUnused
		s.nodes[idx] = t
		return idx
	}

	s.nodes = append(s.nodes, t)
	return len(s.nodes) - 1
}

// removeNode de-allocates a node and makes its slot available for re-use by the IntervalSet
//  If the node is a branch, it de-allocates the entire sub-tree rooted at that branch
func (s *IntervalSet) removeNode(idx int) {
	if idx >= 0 {
		node := s.nodes[idx]
		s.deAllocate(idx)
		if !node.isLeaf {
			s.removeNode(node.left)
			s.removeNode(node.right)
		}
	}
}

// deAllocate de-allocates a single slot and makes it available for re-use by the IntervalSet
func (s *IntervalSet) deAllocate(idx int) {
	if idx >= 0 {
		s.nodeCount--
		s.nodes[idx].nextUnused = s.unused
		s.unused = idx
	}
}

// copyNodes copies a sub-tree from another intervalset to this one
//  it returns in the index in this intervalset of the root of the copied sub-tree
func (s *IntervalSet) copyNodes(other *IntervalSet, root int) int {
	if root == -1 {
		return -1
	}
	newNode := other.nodes[root]
	if !newNode.isLeaf {
		newNode.left = s.copyNodes(other, newNode.left)
		newNode.right = s.copyNodes(other, newNode.right)
	}
	return s.addNode(newNode)
}

func (s *IntervalSet) GetIterator() Iterator {
	s.init()
	return newIterator(s)
}

func NewIntervalSetWithCapacity(cap int) *IntervalSet {
	return &IntervalSet{
		root:   -1,
		unused: -1,
		nodes:  make([]trie, 0, cap),
		ul:     false,
	}
}

func NewIntervalSet(ivals ...interval.Interval) *IntervalSet {
	return NewIntervalSetFromSlice(ivals)
}

func NewIntervalSetFromSlice(ivals []interval.Interval) *IntervalSet {
	set := NewIntervalSetWithCapacity(len(ivals) * 2)
	for _, ival := range ivals {
		set.Add(ival)
	}
	return set
}

func (s *IntervalSet) Clone() *IntervalSet {
	s.init()
	iset := NewIntervalSetWithCapacity(s.Size())
	iset.root = iset.copyNodes(s, s.root)
	iset.ul = s.ul
	return iset
}

func (s *IntervalSet) Add(ival interval.Interval) {
	s.init()
	rhs, rul := trieFromInterval(s, ival)
	s.root = treeOr(s, s, s.ul, rul, s.root, rhs)
	s.ul = s.ul || rul
}

func (s *IntervalSet) Union(other *IntervalSet) *IntervalSet {
	return s.Clone().UnionInPlace(other)
}

func (s *IntervalSet) UnionInPlace(other *IntervalSet) *IntervalSet {
	s.init()
	other.init()
	if other != nil {
		s.root = treeOr(s, other, s.ul, other.ul, s.root, other.root)
		s.ul = s.ul || other.ul
	}
	return s
}

func (s *IntervalSet) Intersection(other *IntervalSet) *IntervalSet {
	if other == nil || other.IsEmpty() {
		return NewIntervalSet()
	}

	return s.Clone().IntersectionInPlace(other)
}

func (s *IntervalSet) IntersectionInPlace(other *IntervalSet) *IntervalSet {
	if s.IsEmpty() {
		return s
	}
	if other == nil || other.IsEmpty() {
		s.removeNode(s.root)
		s.root = -1
		s.ul = false
	} else {
		s.root = treeAnd(s, other, s.ul, other.ul, s.root, other.root)
		s.ul = s.ul && other.ul
	}
	return s
}

func (s *IntervalSet) Difference(other *IntervalSet) *IntervalSet {
	return s.Clone().DifferenceInPlace(other)
}

func (s *IntervalSet) DifferenceInPlace(other *IntervalSet) *IntervalSet {
	s.init()
	other.init()
	if other != nil {
		s.root = treeAnd(s, other, s.ul, !other.ul, s.root, other.root)
		s.ul = s.ul && !other.ul
	}
	return s
}

func (s *IntervalSet) SymmetricDifference(other *IntervalSet) *IntervalSet {
	return s.Clone().SymmetricDifferenceInPlace(other)
}

func (s *IntervalSet) SymmetricDifferenceInPlace(other *IntervalSet) *IntervalSet {
	s.init()
	other.init()
	if other != nil {
		s.root = treeXor(s, other, s.ul, other.ul, s.root, other.root)
		s.ul = s.ul != other.ul
	}
	return s
}

func (s *IntervalSet) Negate() *IntervalSet {
	return s.Clone().NegateInPlace()
}

func (s *IntervalSet) NegateInPlace() *IntervalSet {
	s.init()
	s.ul = !s.ul
	return s
}

func (s *IntervalSet) Extent() interval.Interval {
	s.init()
	if s.root < 0 {
		if s.ul {
			return interval.Unbounded()
		}
		return interval.Empty()
	}
	var (
		lower, upper uint64
		lb, ub       interval.BoundType
	)
	ur := s.ul != s.nodes[s.root].Unbounded()
	if s.ul {
		lower = 0
		lb = interval.Unbound
	} else {
		lower, lb = lowerBound(s, s.root)
	}
	if ur {
		upper = 0
		ub = interval.Unbound
	} else {
		upper, ub = upperBound(s, s.root)
	}
	return interval.NewInterval(lb, lower, upper, ub)
}

func (s *IntervalSet) IntervalContaining(point uint64) (ival interval.Interval) {
	s.init()
	var (
		ok   bool
		left trie
		lir  bool
	)
	ok = s.root >= 0
	if ok {
		left, lir, ok = s.nodes[s.root].leafPreceding(s, point, s.ul)
	}
	if !ok {
		if !s.ul {
			return
		}
		ival = interval.Unbounded()
	} else {
		// See if the left leaf actually represents the left side of an interval
		if lir && (!left.boundBoth() || left.prefix == point) {
			return
		}

		// Set the left edge of the interval appropriately
		switch {
		case left.boundAbove(), left.boundBoth():
			ival = interval.Above(left.prefix)
		case left.boundBelow():
			ival = interval.AtOrAbove(left.prefix)
		}
	}

	// Now go find the right side
	right, rir, ok := s.nodes[s.root].leafSucceeding(s, point, s.ul)
	if !ok {
		return
	} else if !rir {
		// Naw, man, no point
		return interval.Empty()
	}
	// This is the right edge and real and stuff
	switch {
	case right.boundBelow(), right.boundBoth():
		ival = ival.Intersect(interval.Below(right.prefix))
	case right.boundAbove():
		ival = ival.Intersect(interval.AtOrBelow(right.prefix))
	}
	return
}

func (s *IntervalSet) Equals(other *IntervalSet) bool {
	s.init()
	other.init()
	return other != nil && s.ul == other.ul && trieEq(s, other, s.root, other.root)
}

func (s *IntervalSet) Each(f func(interval.Interval)) {
	s.init()
	s.walk(func(lb, ub interval.BoundType, l, u uint64) {
		f(interval.NewInterval(lb, l, u, ub))
	})
}

func (s *IntervalSet) Count() (n int) {
	s.init()
	s.walk(func(_, _ interval.BoundType, _, _ uint64) {
		n++
	})
	return
}

func (s *IntervalSet) Slice() []interval.Interval {
	result := make([]interval.Interval, 0, s.Count())
	s.Each(func(ival interval.Interval) {
		result = append(result, ival)
	})
	return result
}

func (s *IntervalSet) Delta() (total int64) {
	s.Each(func(ival interval.Interval) {
		lower, upper := ival.RightOpenBoundsInt64()
		total += upper - lower
	})
	return
}

func (s *IntervalSet) String() string {
	if s.IsEmpty() {
		return "(Ø)"
	}

	parts := make([]string, 0, s.Count())
	s.Each(func(ival interval.Interval) {
		parts = append(parts, ival.String())
	})
	return strings.Join(parts, ", ")
}

func (s *IntervalSet) IsEmpty() bool {
	if s == nil {
		return true
	}
	s.init()
	return s.root < 0 && !s.ul
}

func (s *IntervalSet) IsUnbounded() bool {
	s.init()
	return s.root < 0 && s.ul
}

func (s *IntervalSet) walk(f func(interval.BoundType, interval.BoundType, uint64, uint64)) {
	if s.root < 0 {
		if s.ul {
			f(interval.Unbound, interval.Unbound, 0, 0)
		}
		return
	}
	var (
		leftkey   uint64
		leftbound interval.BoundType
	)
	do := func(key uint64, left bool, bound interval.BoundType) {
		if left {
			leftkey, leftbound = key, bound
		} else {
			f(leftbound, bound, leftkey, key)
		}
	}
	root := s.nodes[s.root]
	pk, pl, pb := root.iterate(s, do, 0, true, interval.Unbound, s.ul)
	if s.ul != root.Unbounded() {
		do(pk, pl, pb)
		do(0, false, interval.Unbound)
	}
}
