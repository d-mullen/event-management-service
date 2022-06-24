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
	"fmt"
	"sort"
	"strings"

	"github.com/zenoss/zingo/v4/interval"
	"golang.org/x/exp/rand"
)

type (
	value struct {
		value    interface{}
		iset     *IntervalSet
		free     bool
		nextfree int
	}

	trieTemporalSet struct {
		isets    []value
		m        map[interface{}]int
		nextfree int
	}
)

func (lhs *trieTemporalSet) addNew(val interface{}, iset *IntervalSet) int {
	var idx int
	if lhs.nextfree >= 0 {
		idx = lhs.nextfree
		lhs.nextfree = lhs.isets[idx].nextfree
		lhs.isets[idx].value = val
		lhs.isets[idx].iset = iset
		lhs.isets[idx].free = false
	} else {
		lhs.isets = append(lhs.isets, value{
			value:    val,
			iset:     iset,
			nextfree: -1,
		})
		idx = len(lhs.isets) - 1
	}
	lhs.m[val] = idx
	return idx
}

func (lhs *trieTemporalSet) remove(val interface{}) {
	if idx, ok := lhs.m[val]; ok {
		delete(lhs.m, val)
		lhs.isets[idx].free = true
		if lhs.nextfree >= 0 {
			lhs.isets[idx].nextfree = lhs.nextfree
		}
		lhs.nextfree = idx
	}
}

func (lhs *trieTemporalSet) Add(value interface{}, ival interval.Interval) TemporalSet {
	idx, ok := lhs.m[value]
	if !ok {
		lhs.addNew(value, NewIntervalSet(ival))
	} else {
		lhs.isets[idx].iset.Add(ival)
	}
	return lhs
}

func (lhs *trieTemporalSet) AddSlice(values []interface{}, ival interval.Interval) TemporalSet {
	for _, v := range values {
		lhs.Add(v, ival)
	}
	return lhs
}

func (lhs *trieTemporalSet) AddSet(value interface{}, iset *IntervalSet) TemporalSet {
	idx, ok := lhs.m[value]
	if !ok {
		lhs.addNew(value, iset.Clone())
	} else {
		lhs.isets[idx].iset.UnionInPlace(iset)
	}
	return lhs
}

func (lhs *trieTemporalSet) And(rhs TemporalSet) TemporalSet {
	if rhs == nil {
		lhs.nextfree = -1
		lhs.isets = lhs.isets[:0]
		lhs.m = make(map[interface{}]int, len(lhs.m))
		return lhs
	}
	rhst := rhs.(*trieTemporalSet)
	for _, val := range lhs.isets {
		if val.free {
			continue
		}

		isect := val.iset.IntersectionInPlace(rhst.get(val.value))
		if isect.IsEmpty() {
			lhs.remove(val.value)
			continue
		}
	}
	return lhs
}

func (lhs *trieTemporalSet) Or(rhs TemporalSet) TemporalSet {
	if rhs == nil {
		return lhs
	}
	rhst := rhs.(*trieTemporalSet)
	for _, val := range rhst.isets {
		if val.free {
			continue
		}
		idx, ok := lhs.m[val.value]
		if ok {
			lhs.isets[idx].iset.UnionInPlace(val.iset)
		} else {
			lhs.addNew(val.value, val.iset.Clone())
		}
	}
	return lhs
}

func (lhs *trieTemporalSet) Xor(rhs TemporalSet) TemporalSet {
	if rhs == nil {
		return lhs
	}
	rhst := rhs.(*trieTemporalSet)
	for _, val := range rhst.isets {
		if val.free {
			continue
		}

		idx, ok := lhs.m[val.value]
		if ok {
			lhs.isets[idx].iset.SymmetricDifferenceInPlace(val.iset)
		} else {
			lhs.addNew(val.value, val.iset.Clone())
		}
	}
	return lhs
}

func (lhs *trieTemporalSet) AndNot(rhs TemporalSet) TemporalSet {
	if rhs == nil {
		return lhs
	}
	rhst := rhs.(*trieTemporalSet)
	for _, val := range lhs.isets {
		if val.free {
			continue
		}
		symdif := val.iset.DifferenceInPlace(rhst.get(val.value))
		if symdif.IsEmpty() {
			lhs.remove(val.value)
			continue
		}
	}
	return lhs
}

func (set *trieTemporalSet) Negate() TemporalSet {
	for _, val := range set.isets {
		if val.free {
			continue
		}
		val.iset.NegateInPlace()
	}
	return set
}

func (t *trieTemporalSet) Mask(mask *IntervalSet) TemporalSet {
	for _, val := range t.isets {
		if val.free {
			continue
		}
		isect := val.iset.IntersectionInPlace(mask)
		if isect.IsEmpty() {
			t.remove(val.value)
			continue
		}
	}
	return t
}

func (t *trieTemporalSet) MaskInterval(ival interval.Interval) TemporalSet {
	return t.Mask(NewIntervalSet(ival))
}

func (t *trieTemporalSet) Contains(value interface{}, i interval.Interval) bool {
	idx, ok := t.m[value]
	if !ok {
		return false
	}
	return !t.isets[idx].iset.Intersection(NewIntervalSet(i)).IsEmpty()
}

func (t *trieTemporalSet) Cardinality() (n int) {
	for _, v := range t.isets {
		if v.free {
			continue
		}
		if !v.iset.IsEmpty() {
			n++
		}
	}
	return
}

func (t *trieTemporalSet) Empty() bool {
	for _, v := range t.isets {
		if v.free {
			continue
		}
		if !v.iset.IsEmpty() {
			return false
		}
	}
	return true
}

func (t *trieTemporalSet) EachValue(f func(interface{}, []interval.Interval)) {
	var (
		current   interface{}
		intervals []interval.Interval
	)
	t.Each(func(ival interval.Interval, value interface{}) {
		if value != current {
			if current != "" {
				// We did it! Send off a batch
				if len(intervals) > 0 {
					f(current, intervals)
					intervals = make([]interval.Interval, 0)
				}
			}
			current = value
		}
		intervals = append(intervals, ival)
	})
	if len(intervals) > 0 {
		f(current, intervals)
	}
}

func (t *trieTemporalSet) RandValue(ival interval.Interval, src rand.Source) (interface{}, interval.Interval, bool) {
	// Guard against invalid case
	if ival.IsEmpty() {
		return nil, interval.Empty(), false
	}

	// Choose a random point interval
	randts := ival.Rand(src)
	point := interval.Point(randts)

	// Check for impossible intersection before unnecessary masking
	if !t.Extent().Contains(point) {
		return nil, interval.Empty(), false
	}

	// Mask the temporalset to the point
	masked := t.Clone().MaskInterval(point).(*trieTemporalSet)

	// Check for no values at that moment
	if len(masked.m) == 0 {
		return nil, interval.Empty(), false
	}

	// Choose a random map key
	i := rand.New(src).Intn(len(masked.m))
	for _, val := range masked.isets {
		if val.free {
			continue
		}
		if i == 0 {
			return val.value, t.get(val.value).IntervalContaining(randts), true
		}
		i--
	}
	panic("never")

}

func (t *trieTemporalSet) EachValueSet(f func(interface{}, *IntervalSet)) {
	for _, val := range t.isets {
		if val.free {
			continue
		}
		f(val.value, val.iset)
	}
}

func (t *trieTemporalSet) Each(f func(interval.Interval, interface{})) {
	var (
		leftkey   uint64
		leftbound interval.BoundType
		combine   func(uint64, bool, interval.BoundType)
		pk        uint64
		pl        bool
		pb        interval.BoundType
	)
	for _, val := range t.isets {
		if val.free {
			continue
		}

		if val.iset.root < 0 {
			if val.iset.ul {
				f(interval.Unbounded(), val.value)
			} else {
				f(interval.Empty(), val.value)
			}
			continue
		}

		combine = func(key uint64, left bool, bound interval.BoundType) {
			if left {
				leftkey, leftbound = key, bound
			} else {
				ival := interval.NewInterval(leftbound, leftkey, key, bound)
				f(ival, val.value)
			}
		}

		root := val.iset.nodes[val.iset.root]
		pk, pl, pb = root.iterate(val.iset, combine, 0, true, interval.Unbound, val.iset.ul)
		if val.iset.ul != root.Unbounded() {
			combine(pk, pl, pb)
			combine(0, false, interval.Unbound)
		}
	}
}

func (t *trieTemporalSet) Extent() interval.Interval {
	ival := interval.Empty()
	for _, val := range t.isets {
		if val.free {
			continue
		}
		ival = ival.Encompass(val.iset.Extent())
		if ival.IsUnbounded() {
			break
		}
	}
	return ival
}

func (t *trieTemporalSet) Clone() TemporalSet {
	ts := EmptyWithCapacity(len(t.m))
	for _, v := range t.isets {
		if v.free {
			continue
		}
		ts.Update(v.value, v.iset.Clone())
	}
	return ts
}

func (t *trieTemporalSet) ValueSubset(values ...interface{}) TemporalSet {
	tset := EmptyWithCapacity(len(values))
	for _, v := range values {
		if idx, ok := t.m[v]; ok {
			tset.Update(v, t.isets[idx].iset.Clone())
		}
	}
	return tset
}

func (t *trieTemporalSet) AllIntervals() *IntervalSet {
	result := NewIntervalSetWithCapacity(t.Cardinality() * 2)
	for _, val := range t.isets {
		if val.free {
			continue
		}
		result.UnionInPlace(val.iset)
	}
	return result
}

func (t *trieTemporalSet) ContainsValue(value interface{}) bool {
	_, ok := t.m[value]
	return ok
}

func (t *trieTemporalSet) get(value interface{}) *IntervalSet {
	idx, ok := t.m[value]
	if !ok {
		return new(IntervalSet)
	}
	return t.isets[idx].iset
}

func (t *trieTemporalSet) Get(value interface{}) *IntervalSet {
	idx, ok := t.m[value]
	if !ok {
		return NewIntervalSet()
	}
	return t.isets[idx].iset.Clone()
}

func (t *trieTemporalSet) Delete(value interface{}) TemporalSet {
	t.remove(value)
	return t
}

func (t *trieTemporalSet) Update(value interface{}, set *IntervalSet) {
	idx, ok := t.m[value]
	if !ok {
		t.addNew(value, set)
	} else {
		t.isets[idx].iset = set
	}
}

func (t *trieTemporalSet) Values() []interface{} {
	values := make([]interface{}, 0, len(t.isets))
	for _, val := range t.isets {
		if !val.free && !val.iset.IsEmpty() {
			values = append(values, val.value)
		}
	}
	return values
}

func (t *trieTemporalSet) Rand(source rand.Source) (interface{}, interval.Interval) {
	timestamp := t.Extent().Rand(source)
	var (
		p      = interval.Point(timestamp)
		values = make([]interface{}, 0, t.Cardinality())
		ivals  = make([]interval.Interval, 0, t.Cardinality())
	)
	t.EachValue(func(v interface{}, ivs []interval.Interval) {
		for _, i := range ivs {
			if i.Contains(p) {
				values = append(values, v)
				ivals = append(ivals, i)
				return
			} else if p.Before(i) {
				return
			}
		}
	})
	if len(values) == 0 {
		return nil, interval.Empty()
	}
	index := rand.New(source).Intn(len(values))
	return values[index], ivals[index]
}

func (t *trieTemporalSet) MeanValueCount(ival interval.Interval) float64 {
	var sum float64
	t.Clone().MaskInterval(ival).EachInterval(func(iv interval.Interval, vals []interface{}) {
		span := float64(iv.Span())
		sum += float64(len(vals)) * span
	})
	return sum / float64(ival.Span())
}

func (t *trieTemporalSet) String() string {
	var ks []string
	t.EachInterval(func(i interval.Interval, values []interface{}) {
		k := i.String()
		vs := make([]string, len(values))
		for i, v := range values {
			vs[i] = fmt.Sprintf("%v", v)
		}
		sort.Strings(vs)
		ks = append(ks, fmt.Sprintf("%s:{%s}", k, strings.Join(vs, ", ")))
	})
	if len(ks) == 0 {
		return "(Ø)"
	}
	return strings.Join(ks, "; ")
}

func trieFromInterval(s *IntervalSet, ival interval.Interval) (int, bool) {
	var (
		lul, rul    bool
		left, right int
		l, lb       = ival.Lower()
		u, ub       = ival.Upper()
	)
	switch lb {
	case interval.Unbound:
		left = -1
		lul = true
	case interval.EmptyBound:
		left = -1
	case interval.OpenBound:
		left = s.addNode(newLeaf(l, false, true))
	case interval.ClosedBound:
		left = s.addNode(newLeaf(l, true, true))
	}
	switch ub {
	case interval.Unbound:
		right = -1
		rul = true
	case interval.EmptyBound:
		right = -1
	case interval.OpenBound:
		right = s.addNode(newLeaf(u, true, true))
		rul = true
	case interval.ClosedBound:
		right = s.addNode(newLeaf(u, false, true))
		rul = true
	}
	return treeAnd(s, s, lul, rul, left, right), lul && rul
}
