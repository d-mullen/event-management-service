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

	"github.com/zenoss/zingo/v4/interval"
	"golang.org/x/exp/rand"
)

type (
	// TemporalSet models the history of values of an attribute as a collection
	// of pairs <t,v> where t is an interval and v is a string value valid
	// during that interval.
	TemporalSet interface {
		// Add adds the value specified at the interval specified. The current
		// temporal set is returned.
		Add(interface{}, interval.Interval) TemporalSet
		// AddSlice adds all the values specified at the interval specified.
		// The current temporal set is returned.
		AddSlice([]interface{}, interval.Interval) TemporalSet
		// AddSet adds the value specified for the entire interval set
		// specified.  The current temporal set is returned.
		AddSet(interface{}, *IntervalSet) TemporalSet
		// And mutates this set to reflect the intersection with other. The
		// current set is returned.
		And(other TemporalSet) TemporalSet
		// Or mutates this set to reflect the union with the other set. The
		// current set is returned.
		Or(other TemporalSet) TemporalSet
		// Xor mutates this set to reflect the symmetric difference with the
		// other set. The current set is returned.
		Xor(other TemporalSet) TemporalSet
		// AndNot mutates this set to reflect the intersection with the
		// negation of the other set. The current set is returned.
		AndNot(other TemporalSet) TemporalSet
		// Negate mutates this set to reflect its negation, and returns it.
		Negate() TemporalSet
		// Mask eliminates intervals for all values that do not fall within the
		// intervals specified.
		Mask(*IntervalSet) TemporalSet
		// MaskInterval masks within a single interval
		MaskInterval(interval.Interval) TemporalSet
		// Each executes the callback provided for each temporal atom in the
		// set
		Each(func(interval.Interval, interface{}))
		// EachValue executes the callback provided for each value in the set
		EachValue(func(interface{}, []interval.Interval))
		// EachValueSet executes the callback provided for each value in the set
		EachValueSet(func(interface{}, *IntervalSet))
		// EachInterval executes the callback provided for each distinct
		// interval in the set with the values valid during that interval.
		EachInterval(func(interval.Interval, []interface{}))
		// Contains returns whether the given ids and intervals exist within
		// this set
		Contains(interface{}, interval.Interval) bool
		// Empty returns whether this temporal set represents any values
		Empty() bool
		// Cardinality returns the number of values represented in this set
		// over at least one non-empty interval.
		Cardinality() int
		// Extent returns the smallest interval encompassing all
		// temporal atoms in this set.
		Extent() interval.Interval
		// Clone produces a temporal set identical to this one that will be
		// unaffected by updates to the original.
		Clone() TemporalSet
		// ValueSubset produces a subset of this temporal set containing only
		// those values specified.
		ValueSubset(values ...interface{}) TemporalSet
		// AllIntervals returns an interval set containing the intervals
		// covered by any value in this set
		AllIntervals() *IntervalSet
		// ContainsValue returns whether the value is represented in the set
		// for any interval
		ContainsValue(interface{}) bool
		// Get returns the interval set corresponding to the intervals for
		// which the value specified was valid.
		Get(interface{}) *IntervalSet
		// Delete removes the specified value from the temporalset
		Delete(interface{}) TemporalSet
		// Update sets the interval set for a value to the one provided
		Update(interface{}, *IntervalSet)
		// Values returns all the unique values in the the temporal set
		Values() []interface{}
		// RandValue chooses a random value from those active during a random
		// point within the interval provided, and a boolean indicating whether
		// any values were active at the point chosen.
		RandValue(interval.Interval, rand.Source) (interface{}, interval.Interval, bool)
		// MeanValueCount returns the mean number of concurrent values across
		// the interval specified
		MeanValueCount(interval.Interval) float64

		fmt.Stringer
	}
)

var (
	// Above returns a temporal set comprising the temporal atom
	// (instant, ∞):{value}
	Above = unboundedTrieSetFactory(false, false, true)
	// AtOrAbove returns a temporal set comprising the temporal atom
	// [instant, ∞):{value}
	AtOrAbove = unboundedTrieSetFactory(false, true, true)
	// Below returns a temporal set comprising the temporal atom
	// (-∞, instant):{value}
	Below = unboundedTrieSetFactory(true, true, true)
	// AtOrBelow returns a temporal set comprising the temporal atom
	// (-∞, instant]:{value}
	AtOrBelow = unboundedTrieSetFactory(true, false, true)
	// Hole returns a temporal set comprising the temporal atoms
	// (-∞, instant):{value}; (instant, ∞):{value}
	Hole = unboundedTrieSetFactory(true, true, false)
	// Point returns a temporal set comprising the temporal atom
	// [instant, instant]:{value}
	Point = unboundedTrieSetFactory(false, true, false)
	// Open returns a temporal set comprising the temporal atom
	// (lower, upper):{values...}
	Open = boundedTrieFactory(interval.Open)
	// Closed returns a temporal set comprising the temporal atom
	// [lower, upper]:{values...}
	Closed = boundedTrieFactory(interval.Closed)
	// LeftOpen returns a temporal set comprising the temporal atom
	// (lower, upper]:{values...}
	LeftOpen = boundedTrieFactory(interval.LeftOpen)
	// RightOpen returns a temporal set comprising the temporal atom
	// [lower, upper):{values...}
	RightOpen = boundedTrieFactory(interval.RightOpen)
)

// FromInterval returns a temporal set bounded by the interval specified,
// containing the specified value
func FromInterval(i interval.Interval, value interface{}) TemporalSet {
	tset := Empty()
	tset.AddSet(value, NewIntervalSet(i))
	return tset
}

// Empty returns an empty temporal set, (Ø)
func Empty() TemporalSet {
	return EmptyWithCapacity(1)
}

func EmptyWithCapacity(cap int) TemporalSet {
	return &trieTemporalSet{
		isets:    make([]value, 0, cap),
		nextfree: -1,
		m:        make(map[interface{}]int, cap),
	}
}

func unboundedTrieSetFactory(treeUnboundedLeft, includesValue, unbounded bool) func(uint64, interface{}) TemporalSet {
	return func(moment uint64, value interface{}) TemporalSet {
		tset := Empty()
		iset := NewIntervalSetWithCapacity(5)
		iset.root = iset.addNode(newLeaf(moment, includesValue, unbounded))
		iset.ul = treeUnboundedLeft
		tset.(*trieTemporalSet).addNew(value, iset)
		return tset
	}
}

func boundedTrieFactory(ivalfactory func(uint64, uint64) interval.Interval) func(uint64, uint64, []interface{}) TemporalSet {
	return func(lower, upper uint64, values []interface{}) TemporalSet {
		tset := EmptyWithCapacity(len(values))
		iset := NewIntervalSet(ivalfactory(lower, upper))
		for _, m := range values {
			tset.AddSet(m, iset.Clone())
		}
		return tset
	}
}
