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

import "github.com/zenoss/zingo/v4/interval"

// Iterator allows iteration over an IntervalSet.
// Each state of the iterator represents an edge up/down transition point
// of the IntervalSet with attributes t,s,b, where
// 	t is the time at which the transition occurs;
//	s is the boolean state for times after the transition;
// 	b indicates whether the transition includes t
// Typical usage:
// for it := iset.GetIterator(); !it.Done(); it.Next() {
//     time_value, state, bound_type = it.Value()
// }
type Iterator struct {
	stack []int
	state bool
	both  boundBothState
	iset  *IntervalSet
}

// boundBothState represents the state of the node on the top of the stack
// for which leaf.boundBoth() == true.  These nodes represent two transitions;
// the first, which embodies the normal sense of the boundary, and the second,
// with an inverted boundary.
type boundBothState int8

const (
	irrelevant boundBothState = iota
	first
	second
)

// Create a new Iterator given a trie and a bool indicating whether the
// trie is to be considered left unbounded.
func newIterator(s *IntervalSet) Iterator {
	if s.root < 0 {
		return Iterator{}
	}
	it := Iterator{
		stack: make([]int, 0, s.nodes[s.root].Level()),
		state: !s.ul,
		both:  0,
		iset:  s,
	}
	it.push(s.root)
	return it
}

// Value returns the value, state, and bound type of the iterator.
func (it *Iterator) Value() (uint64, bool, interval.BoundType) {
	if it.Done() {
		return 0, false, interval.EmptyBound
	}
	// Return the state at the top of the stack.
	top := it.stack[len(it.stack)-1]
	l := it.iset.nodes[top]
	bound := interval.OpenBound
	if (it.state == l.includesValue) == (it.both != second) {
		bound = interval.ClosedBound
	}
	return l.prefix, it.state, bound
}

// Next increments the iterator to its next value
func (it *Iterator) Next() {
	if it.Done() {
		return
	}
	// Invert the in/out sense of the iterator
	it.state = !it.state
	// If the leaf node on the top of the stack is a "boundBoth",
	// leave it on the stack but decrement the boundBoth state
	if it.both == first {
		it.both = second
		return
	}
	// Pop the current node off the top of the stack.
	it.stack = it.stack[:len(it.stack)-1]
	if it.Done() {
		return
	}
	// Advance the iterator to its next state
	top := it.iset.nodes[it.stack[len(it.stack)-1]]
	if !top.isLeaf {
		// Replace the branch node by its subtrees
		it.stack = it.stack[:len(it.stack)-1]
		it.push(top.right)
	}
}

// Done returns true if the iterator has reached its end
func (it *Iterator) Done() bool {
	return len(it.stack) == 0
}

// push appends a subtree onto an iterator
func (it *Iterator) push(n int) {
	// Walk down the left hand branch of the tree until we hit a leaf, pushing
	// each branch node onto the stack for subsequent traversal of its right hand child.
	for b := it.iset.nodes[n]; !b.isLeaf; b = it.iset.nodes[n] {
		it.stack = append(it.stack, n)
		n = b.left
	}
	//  We've reached a leaf node, which represents the iterator's next state.
	//  Push it onto the stack; if it is a bound-both node, flag the state as such.
	it.stack = append(it.stack, n)
	if it.iset.nodes[n].boundBoth() {
		it.both = first
	} else {
		it.both = irrelevant
	}
}
