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

func (l trie) boundBelow() bool {
	return l.includesValue && l.unbounded
}

func (l trie) boundAbove() bool {
	return !l.includesValue && l.unbounded
}

func (l trie) boundBoth() bool {
	return l.includesValue && !l.unbounded
}

func (l trie) leafLower() (uint64, interval.BoundType) {
	if l.boundBelow() || l.boundBoth() {
		return l.prefix, interval.ClosedBound
	}
	return l.prefix, interval.OpenBound
}

func (l trie) leafUpper() (uint64, interval.BoundType) {
	if l.boundAbove() || l.boundBoth() {
		return l.prefix, interval.ClosedBound
	}
	return l.prefix, interval.OpenBound
}

func (l trie) leafLeafPreceding(key uint64, ul bool) (trie, bool, bool) {
	switch {
	case l.prefix < key:
		return l, ul, true
	case l.prefix == key && !l.boundAbove():
		return l, ul, true
	}
	return trie{}, false, false
}

func (l trie) leafLeafSucceeding(key uint64, ul bool) (trie, bool, bool) {
	switch {
	case l.prefix > key:
		return l, ul, true
	case l.prefix == key && !l.boundBelow():
		return l, ul, true
	}
	return trie{}, false, false
}

func (t trie) leafIterate(f func(key uint64, left bool, bound interval.BoundType), prevkey uint64, prevleft bool, prevbound interval.BoundType, unboundedLeft bool) (uint64, bool, interval.BoundType) {
	key := t.prefix
	if t.boundBelow() {
		if unboundedLeft {
			f(prevkey, prevleft, prevbound)
			f(key, false, interval.OpenBound)
		}
		return key, true, interval.ClosedBound
	} else if t.boundAbove() {
		if unboundedLeft {
			f(prevkey, prevleft, prevbound)
			f(key, false, interval.ClosedBound)
		}
		return key, true, interval.OpenBound
	} else if t.boundBoth() {
		if unboundedLeft {
			f(prevkey, prevleft, prevbound)
			f(key, false, interval.OpenBound)
		} else {
			f(key, true, interval.ClosedBound)
			f(key, false, interval.ClosedBound)
		}
		return key, true, interval.OpenBound
	}
	return key, true, interval.Unbound
}

func newLeaf(prefix uint64, includesValue, unbounded bool) trie {
	return trie{
		prefix:        prefix,
		level:         0,
		unbounded:     unbounded,
		isLeaf:        true,
		includesValue: includesValue,
		left:          -1,
		right:         -1,
		nextUnused:    -1,
	}
}

func maybeNewLeaf(lset, rset *IntervalSet, boundBelow, boundAbove bool, aidx, bidx int) int {
	a, b := lset.nodes[aidx], rset.nodes[bidx]
	includesValue := boundBelow
	unbounded := boundBelow != boundAbove
	if !boundBelow && !boundAbove {
		// De-allocate a & b
		lset.removeNode(aidx)
		if lset == rset {
			lset.removeNode(bidx)
		}
		return -1
	}
	if includesValue == a.includesValue && unbounded == a.unbounded {
		// De-allocate b, return a
		if lset == rset {
			lset.removeNode(bidx)
		}
		return aidx
	}
	if includesValue == b.includesValue && unbounded == b.unbounded {
		// De-allocate a, return b
		lset.removeNode(aidx)
		if lset != rset {
			// Copy node from rset to lset
			bidx = lset.copyNodes(rset, bidx)
		}
		return bidx
	}
	// De-allocate a and b and add new leaf
	lset.removeNode(aidx)
	if lset == rset {
		lset.removeNode(bidx)
	}
	return lset.addNode(newLeaf(a.prefix, includesValue, unbounded))
}
