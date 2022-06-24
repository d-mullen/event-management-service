# temporalset
--
    import "github.com/zenoss/zingo/v4/temporalset"


## Usage

```go
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
```

#### type IntervalSet

```go
type IntervalSet struct {
}
```


#### func  NewIntervalSet

```go
func NewIntervalSet(ivals ...interval.Interval) *IntervalSet
```

#### func  NewIntervalSetFromSlice

```go
func NewIntervalSetFromSlice(ivals []interval.Interval) *IntervalSet
```

#### func (*IntervalSet) Add

```go
func (s *IntervalSet) Add(ival interval.Interval)
```

#### func (*IntervalSet) Clone

```go
func (s *IntervalSet) Clone() *IntervalSet
```

#### func (*IntervalSet) Count

```go
func (s *IntervalSet) Count() (n int)
```

#### func (*IntervalSet) Difference

```go
func (s *IntervalSet) Difference(other *IntervalSet) *IntervalSet
```

#### func (*IntervalSet) Each

```go
func (s *IntervalSet) Each(f func(interval.Interval))
```

#### func (*IntervalSet) Equals

```go
func (s *IntervalSet) Equals(other *IntervalSet) bool
```

#### func (*IntervalSet) Extent

```go
func (s *IntervalSet) Extent() interval.Interval
```

#### func (*IntervalSet) Intersection

```go
func (s *IntervalSet) Intersection(other *IntervalSet) *IntervalSet
```

#### func (*IntervalSet) IsEmpty

```go
func (s *IntervalSet) IsEmpty() bool
```

#### func (*IntervalSet) IsUnbounded

```go
func (s *IntervalSet) IsUnbounded() bool
```

#### func (*IntervalSet) Negate

```go
func (s *IntervalSet) Negate() *IntervalSet
```

#### func (*IntervalSet) Op

```go
func (s *IntervalSet) Op(op SetOperation, other *IntervalSet) *IntervalSet
```
Op is a testing shortcut. It accepts one of the set operations &, |, ^, - and
performs the appropriate operation on the given interval set.

#### func (*IntervalSet) Slice

```go
func (s *IntervalSet) Slice() []interval.Interval
```

#### func (*IntervalSet) String

```go
func (s *IntervalSet) String() string
```

#### func (*IntervalSet) SymmetricDifference

```go
func (s *IntervalSet) SymmetricDifference(other *IntervalSet) *IntervalSet
```

#### func (*IntervalSet) Union

```go
func (s *IntervalSet) Union(other *IntervalSet) *IntervalSet
```

#### type SetOperation

```go
type SetOperation string
```


```go
const (
	OpIntersection        SetOperation = "&"
	OpUnion               SetOperation = "|"
	OpSymmetricDifference SetOperation = "^"
	OpDifference          SetOperation = "-"
)
```

#### type TemporalSet

```go
type TemporalSet interface {
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
	// Op is a testing shortcut. It accepts one of the strings &, |, ^, -
	// and performs the appropriate operation on the given set.
	Op(SetOperation, TemporalSet) TemporalSet
	// Mask eliminates intervals for all values that do not fall within the
	// intervals specified.
	Mask(...interval.Interval) TemporalSet
	// Each executes the callback provided for each temporal atom in the
	// set
	Each(func(interval.Interval, interface{}))
	// EachValue executes the callback provided for each value in the set
	EachValue(func(interface{}, []interval.Interval))
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

	fmt.Stringer
}
```

TemporalSet models the history of values of an attribute as a collection of
pairs <t,v> where t is an interval and v is a string value valid during that
interval.

#### func  Empty

```go
func Empty() TemporalSet
```
Empty returns an empty temporal set, (Ø)

#### func  EmptyWithCapacity

```go
func EmptyWithCapacity(cap int) TemporalSet
```

#### func  FromInterval

```go
func FromInterval(i interval.Interval, value interface{}) TemporalSet
```
FromInterval returns a temporal set bounded by the interval specified,
containing the specified value
