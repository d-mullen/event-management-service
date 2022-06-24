# interval
--
    import "github.com/zenoss/zingo/v4/temporalset/interval"


## Usage

#### type Bound

```go
type Bound uint8
```


```go
const (
	EmptyBound Bound = iota
	Unbound
	OpenBound
	ClosedBound
)
```

#### type Interval

```go
type Interval interface {
	fmt.Stringer

	// Lower returns the lower bound of the interval and whether it is open or
	// closed
	Lower() (uint64, Bound)

	// Upper returns the upper bound of the interval and whether it is open or
	// closed
	Upper() (uint64, Bound)

	// Empty returns true if the interval is empty
	Empty() bool

	// Unbounded returns true if the interval is unbounded
	Unbounded() bool

	// Equals returns true if the intervals have the same bounds
	Equals(Interval) bool

	// Intersect returns the intersection of another interval with this one.
	Intersect(Interval) Interval
}
```

Interval represents a continuous range of integers

#### func  ClosedInterval

```go
func ClosedInterval(lower, upper uint64) Interval
```

#### func  EmptyInterval

```go
func EmptyInterval() Interval
```

#### func  IntervalAbove

```go
func IntervalAbove(lower uint64) Interval
```

#### func  IntervalAtOrAbove

```go
func IntervalAtOrAbove(lower uint64) Interval
```

#### func  IntervalAtOrBelow

```go
func IntervalAtOrBelow(upper uint64) Interval
```

#### func  IntervalBelow

```go
func IntervalBelow(upper uint64) Interval
```

#### func  LeftOpenInterval

```go
func LeftOpenInterval(lower, upper uint64) Interval
```

#### func  NewInterval

```go
func NewInterval(lowerBound Bound, lower, upper uint64, upperBound Bound) Interval
```

#### func  OpenInterval

```go
func OpenInterval(lower, upper uint64) Interval
```

#### func  RightOpenInterval

```go
func RightOpenInterval(lower, upper uint64) Interval
```

#### func  UnboundedInterval

```go
func UnboundedInterval() Interval
```
