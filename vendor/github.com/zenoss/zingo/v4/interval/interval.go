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

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/pkg/errors"
	"golang.org/x/exp/rand"
)

type BoundType uint8

const (
	EmptyBound BoundType = iota
	Unbound
	ClosedBound
	OpenBound
)

// Interval represents a continuous range of integers
type Interval struct {
	lower      uint64
	upper      uint64
	lowerBound BoundType
	upperBound BoundType
}

func Complement(a BoundType) BoundType {
	switch a {
	case ClosedBound:
		return OpenBound
	case OpenBound:
		return ClosedBound
	case Unbound:
		return EmptyBound
	case EmptyBound:
		return Unbound
	}
	panic("Unknown bound type")
}

func NewInterval(lowerBound BoundType, lower, upper uint64, upperBound BoundType) Interval {
	return Interval{
		lower:      lower,
		upper:      upper,
		lowerBound: lowerBound,
		upperBound: upperBound,
	}
}

func Open(lower, upper uint64) Interval {
	return NewInterval(OpenBound, lower, upper, OpenBound)
}

func Closed(lower, upper uint64) Interval {
	return NewInterval(ClosedBound, lower, upper, ClosedBound)
}

func LeftOpen(lower, upper uint64) Interval {
	return NewInterval(leftOpenBounds(lower, upper))
}

func RightOpen(lower, upper uint64) Interval {
	return NewInterval(rightOpenBounds(lower, upper))
}

func Empty() Interval {
	return NewInterval(emptyBounds())
}

func Unbounded() Interval {
	return NewInterval(unboundedBounds())
}

func Below(upper uint64) Interval {
	return NewInterval(belowBounds(upper))
}

func AtOrBelow(upper uint64) Interval {
	return NewInterval(atOrBelowBounds(upper))
}

func Above(lower uint64) Interval {
	return NewInterval(aboveBounds(lower))
}

func AtOrAbove(lower uint64) Interval {
	return NewInterval(atOrAboveBounds(lower))
}

func Point(i uint64) Interval {
	return NewInterval(pointBounds(i))
}

// Lower returns the lower bound of an interval and whether it is open or
// closed
func (i Interval) Lower() (uint64, BoundType) {
	return i.lower, i.lowerBound
}

// Upper returns the upper bound of an interval and whether it is open or
// closed
func (i Interval) Upper() (uint64, BoundType) {
	return i.upper, i.upperBound
}

func (i Interval) String() string {
	if i.IsEmpty() {
		return "(Ø)"
	}
	var lb, ub, lv, uv string
	lv, uv = fmt.Sprintf("%d", i.lower), fmt.Sprintf("%d", i.upper)
	switch i.lowerBound {
	case OpenBound:
		lb = "("
	case ClosedBound:
		lb = "["
	case Unbound:
		lb = "("
		lv = "-∞"
	default:
		// This shouldn't be possible because of the i.Empty() check above
		lb = "("
	}
	switch i.upperBound {
	case OpenBound:
		ub = ")"
	case ClosedBound:
		ub = "]"
	case Unbound:
		ub = ")"
		uv = "∞"
	default:
		// This shouldn't be possible because of the i.Empty() check above
		ub = ")"
	}
	return fmt.Sprintf("%s%s, %s%s", lb, lv, uv, ub)
}

// Span returns the total size of the interval
func (i Interval) Span() uint64 {
	if i.IsEmpty() {
		return 0
	}
	if i.IsHalfUnbounded() {
		return math.MaxUint64
	}
	l, u := i.NormalizedBounds()
	return u - l
}

func (i Interval) Rand(source rand.Source) uint64 {
	l, u := i.NormalizedBounds()
	return rand.New(source).Uint64n(u-l) + l
}

// IsEmpty returns true if the interval is empty
func (i Interval) IsEmpty() bool {
	return isEmpty(i.lowerBound, i.lower, i.upper, i.upperBound)
}

// IsUnbounded returns true if the interval is unbounded on both sides
func (i Interval) IsUnbounded() bool {
	return isUnbounded(i.lowerBound, i.upperBound)
}

// IsHalfUnbounded returns true if the interval is unbounded on either side
func (i Interval) IsHalfUnbounded() bool {
	return isHalfUnbounded(i.lowerBound, i.upperBound)
}

// Equals returns true if the intervals have the same bounds
func (i Interval) Equals(other Interval) bool {
	if i.IsEmpty() {
		return isEmpty(other.Bounds())
	}
	if isEmpty(other.Bounds()) {
		return false
	}

	l, lb := other.Lower()
	u, ub := other.Upper()
	return boundsEqual(i.lower, i.lowerBound, l, lb) && boundsEqual(i.upper, i.upperBound, u, ub)
}

// LessThan returns true if the original interval starts before the given
// interval, or if the lower bounds are equal, if the original interval
// ends before the given interval.
func (i Interval) LessThan(other Interval) bool {
	olb, olbt := other.Lower()
	if hasLowerLowerBound(i.lower, i.lowerBound, olb, olbt) {
		return true
	}

	al, alb := i.lower, i.lowerBound
	bl, blb := other.Lower()

	if al == bl && alb == blb {
		oub, oubt := other.Upper()
		return hasLowerUpperBound(i.upper, i.upperBound, oub, oubt)
	}

	return false
}

// Bounds returns the constituent parts of the interval
func (i Interval) Bounds() (BoundType, uint64, uint64, BoundType) {
	return i.lowerBound, i.lower, i.upper, i.upperBound
}

// NormalizedBounds returns the bounds of a congruent right-open interval.
func (i Interval) NormalizedBounds() (l, u uint64) {
	if i.IsEmpty() {
		return
	}
	switch i.lowerBound {
	case Unbound:
		l = 0
	case ClosedBound:
		l = i.lower
	case OpenBound:
		l = i.lower + 1
	}
	switch i.upperBound {
	case Unbound:
		u = math.MaxUint64
	case ClosedBound:
		u = i.upper + 1
	case OpenBound:
		u = i.upper
	}
	return
}

// Contains returns true if the original interval contains the given
// interval completely.
// An Empty interval can only contain another empty interval.
// An unbounded interval can only be contained by another unbounded interval.
func (i Interval) Contains(other Interval) bool {
	olbt, olb, oub, oubt := other.Bounds()
	return contains(
		i.lower,
		i.lowerBound,
		i.upper,
		i.upperBound,
		olb,
		olbt,
		oub,
		oubt,
	)
}

// Before returns true if the original interval is completely before the
// given interval.  If either interval is empty, Before ALWAYS returns false
func (i Interval) Before(other Interval) bool {
	if i.IsEmpty() || isEmpty(other.Bounds()) {
		return false
	}

	l, lb := other.Lower()
	if i.upperBound == Unbound || lb == Unbound {
		return false
	}
	return !((i.upper > l) || i.upper == l && i.upperBound == ClosedBound && lb == ClosedBound)
}

// Intersect returns the intersection of another interval with this one.
func (i Interval) Intersect(other Interval) Interval {
	olbt, olb, oub, oubt := other.Bounds()
	return NewInterval(intersection(
		i.lower,
		i.lowerBound,
		i.upper,
		i.upperBound,
		olb,
		olbt,
		oub,
		oubt,
	))
}

// Partition returns three intervals: left difference, intersection, right
// difference
func (i Interval) Partition(other Interval) (Interval, Interval, Interval) {
	var (
		sm = i
		lg = other
	)
	if lg.LessThan(sm) {
		sm, lg = lg, sm
	}
	isect := sm.Intersect(lg)
	if isect.IsEmpty() {
		return sm, isect, lg
	}
	il, ilb := isect.Lower()
	iu, iub := isect.Upper()
	sl, slb := sm.Lower()
	var (
		lu  uint64
		lub BoundType
	)
	if sm.Contains(lg) {
		lu, lub = sm.Upper()
	} else {
		lu, lub = lg.Upper()
	}
	return NewInterval(slb, sl, il, Complement(ilb)), isect, NewInterval(Complement(iub), iu, lu, lub)
}

// Encompass returns the minimal interval that encompasses both.
func (i Interval) Encompass(other Interval) Interval {
	olbt, olb, oub, oubt := other.Bounds()
	return NewInterval(encompass(
		i.lower,
		i.lowerBound,
		i.upper,
		i.upperBound,
		olb,
		olbt,
		oub,
		oubt,
	))
}

func (i Interval) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"l":  i.lower,
		"lb": i.lowerBound,
		"u":  i.upper,
		"ub": i.upperBound,
	})
}

func (i *Interval) UnmarshalJSON(data []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber() // need this or we can't handle very large numbers
	r := make(map[string]interface{})
	if err := decoder.Decode(&r); err != nil {
		return err
	}
	lb, err := strconv.ParseUint(r["lb"].(json.Number).String(), 10, 64)
	if err != nil {
		return errors.Wrap(err, "error unmarshalling lower bound type")
	}
	i.lowerBound = BoundType(lb)

	l, err := strconv.ParseUint(r["l"].(json.Number).String(), 10, 64)
	if err != nil {
		return errors.Wrap(err, "error unmarshalling lower bound")
	}
	i.lower = uint64(l)

	u, err := strconv.ParseUint(r["u"].(json.Number).String(), 10, 64)
	if err != nil {
		return errors.Wrap(err, "error unmarshalling upper bound")
	}
	i.upper = uint64(u)

	ub, err := strconv.ParseUint(r["ub"].(json.Number).String(), 10, 64)
	if err != nil {
		return errors.Wrap(err, "error unmarshalling upper bound type")
	}
	i.upperBound = BoundType(ub)

	return nil
}

// RightOpenBoundsInt64 returns the bounds as int64s: [lower, upper)
func (i Interval) RightOpenBoundsInt64() (lower, upper int64) {
	l, u := i.NormalizedBounds()
	lower, upper = int64(l), int64(u)
	if lower < 0 {
		lower = math.MaxInt64
	}
	if upper < 0 {
		upper = math.MaxInt64
	}
	return
}

// LeftOpenBoundsInt64 returns the bounds as int64s: (lower, upper]
//  If the lower bound is unbounded or closed on 0, lower will be -1
func (i Interval) LeftOpenBoundsInt64() (lower, upper int64) {
	lt, l, u, ut := i.Bounds()
	lower, upper = int64(l), int64(u)
	switch {
	case lt == Unbound:
		lower = -1
	case lower < 0:
		lower = math.MaxInt64
	case lt == ClosedBound:
		lower--
	}

	switch {
	case ut == Unbound:
		upper = math.MaxInt64
	case upper < 0:
		upper = math.MaxInt64
	case ut == OpenBound:
		upper--
	}
	return
}

func UnmarshalIntervalArray(data []byte) ([]Interval, error) {
	result := make([]Interval, 0)
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, errors.WithStack(err)
	}
	ivals := make([]Interval, len(result))
	for i, ival := range result {
		ivals[i] = ival
	}
	return ivals, nil
}

func outermost(a, b BoundType) BoundType {
	if a < b {
		return a
	}
	return b
}
