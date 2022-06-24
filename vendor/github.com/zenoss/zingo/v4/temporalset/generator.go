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
	"container/heap"

	"github.com/zenoss/zingo/v4/interval"
)

// EachInterval executes the callback provided for each distinct interval in the set with
// the values valid during that interval.  This routine is in a separate file in order to
// isolate its helper classes and functions from the main implementation of the class.
func (t *trieTemporalSet) EachInterval(f func(interval.Interval, []interface{})) {
	// active is a list of keys for which the corresponding iterator is in a "true" state;
	//   it is the list with which f will be called.
	active := make([]interface{}, 0, len(t.isets))

	// backPtr is a list parallel to the active list, which allows us to update the index in
	//   the corresponding keyIteration when the keys changes position in the active list.
	//   A nil backPtr indicates the iterator for that key is no longer active
	backPtr := make([]*keyIteration, 0, len(t.isets))

	// queue is a priority queue of pointers to keyIteration structs.  Presence of a key
	// in the queue indicates that the key's iterator is not exhausted.
	queue := make(keyIterationPriorityQueue, 0, len(t.isets))

	// Initialize data structures for traversal
	for _, val := range t.isets {
		if val.free {
			continue
		}
		k, v := val.value, val.iset
		if v.IsEmpty() {
			// This key is never in a "true" state.  Ignore it.
			continue
		} else if v.IsUnbounded() {
			// This key is always in a "true" state.  Add it to the active list,
			// but since it never changes state it doesn't need to be in the heap.
			active = append(active, k)
			backPtr = append(backPtr, nil)
			continue
		}
		// This key will be changing state.  Add it to the heap.
		it := v.GetIterator()
		t, s, b := it.Value()
		vb := keyIteration{k, it, getHeapPriority(t, s, b), -1}
		queue = append(queue, &vb)
		// If the first transition in the iterator is to a "false" state, its initial
		// state must be true.  Add it to the active list.
		if s == false {
			vb.idx = len(active)
			active = append(active, k)
			backPtr = append(backPtr, &vb)
		}
	}
	heap.Init(&queue)

	lowerVal := uint64(0)
	lowerBound := interval.Unbound
	for len(queue) > 0 {
		// Grab the keyIteration off the top of the heap.  This will have the lowest key.
		top := queue[0]

		// Close off the active interval and perform the callback.
		t, s, b := top.it.Value()
		b = toUpper(s, b)
		if len(active) > 0 {
			intv := interval.NewInterval(lowerBound, lowerVal, t, b)
			f(intv, active)
		}
		// Update the beginning of the next interval
		lowerVal, lowerBound = t, interval.Complement(b)

		// Update the active list for the next callback.  Update all keys with the
		// same transition point as the one on top (including the one on top).
		priority := top.priority
		for len(queue) > 0 && queue[0].priority == priority {
			top = queue[0]
			_, s, _ := top.it.Value()
			if s {
				// Add it to the active list
				top.idx = len(active)
				active = append(active, top.key)
				backPtr = append(backPtr, top)
			} else {
				// Remove it from the active list.
				idx := top.idx
				// In the active/backpointer lists, swap the values with those at the end of the list
				// so that we can resize the list.  If they're already at the end, don't bother.
				if idx != len(active)-1 {
					active[idx], active[len(active)-1] = active[len(active)-1], active[idx]
					backPtr[idx], backPtr[len(backPtr)-1] = backPtr[len(backPtr)-1], backPtr[idx]
					if backPtr[idx] != nil {
						backPtr[idx].idx = idx
					}
				}
				// Resize the lists
				active = active[:len(active)-1]
				backPtr = backPtr[:len(backPtr)-1]
				// Update the keyIteration to indicate it is not in the active list
				top.idx = -1
			}
			// Increment the iterator
			top.it.Next()
			if top.it.Done() {
				// This keys's iterator is done; its state will never change again.
				// Remove it from the heap
				if top.idx != -1 {
					backPtr[top.idx] = nil
				}
				heap.Pop(&queue)
			} else {
				// Update the key's priority.
				t, s, b := top.it.Value()
				top.priority = getHeapPriority(t, s, b)
				heap.Fix(&queue, 0)
			}
		}
	}

	// Perform final callback for remaining
	if len(active) > 0 {
		intv := interval.NewInterval(lowerBound, lowerVal, 0, interval.Unbound)
		f(intv, active)
	}
}

// heapPriority is an combination of the time and boundtype for easy comparison.
type heapPriority struct {
	val    uint64
	bigger bool
}

func getHeapPriority(t uint64, s bool, b interval.BoundType) heapPriority {
	bound := s == (b == interval.OpenBound)
	return heapPriority{t, bound}
}

func toUpper(s bool, b interval.BoundType) interval.BoundType {
	if s {
		b = interval.Complement(b)
	}
	return b
}

// keyIteration represents the state of an IntervalSet iteration for a given key.  It contains additional
// information for use in the priority queue and coordination with the list of active keys.
type keyIteration struct {
	// "key" is the identifier for this struct.  It is the key in the map in the original temporalSet.
	key interface{}
	// "it" is an iterator over the IntervalSet for this value.
	it Iterator
	// "priority" is a representation of the iterator's current state for fast comparision
	priority heapPriority
	// "idx" is the index into the "active" list corresponding to this object.  A value of -1
	// implies that the key is not in the active list.
	idx int
}

// keyIterationPriorityQueue is a list of pointers to keyIteration structures.  It is kept in heap
// order and accessed as a priority queue ordered by the key's heapPriority.
type keyIterationPriorityQueue []*keyIteration

func (h keyIterationPriorityQueue) Len() int      { return len(h) }
func (h keyIterationPriorityQueue) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h keyIterationPriorityQueue) Less(i, j int) bool {
	lhs, rhs := h[i].priority, h[j].priority
	if lhs.val != rhs.val {
		return lhs.val < rhs.val
	}
	return !lhs.bigger && rhs.bigger
}

func (h *keyIterationPriorityQueue) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*keyIteration))
}

func (h *keyIterationPriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
