package frequency

import (
	"math"
	"sync"
	"sync/atomic"

	"github.com/zenoss/zingo/v4/interval"
	"github.com/zenoss/zingo/v4/temporalset"
)

type Bucket struct {
	Key    map[string]any
	Values []int64
}

type FrequencyMap struct {
	mu         sync.RWMutex
	buckets    map[string]*Bucket
	maxBuckets int

	groupBy       []string
	startTime     int64
	endTime       int64
	downsample    int64
	persistCounts bool
}

func NewFrequencyMap(groupBy []string, startTime, endTime, downsample int64, persistCounts bool) *FrequencyMap {
	return &FrequencyMap{
		buckets:       make(map[string]*Bucket),
		groupBy:       groupBy,
		startTime:     startTime,
		endTime:       endTime,
		downsample:    downsample,
		persistCounts: persistCounts,
	}
}

func (f *FrequencyMap) GroupBy() []string {
	results := make([]string, len(f.groupBy))
	copy(results, f.groupBy)
	return results
}

func (f *FrequencyMap) Get() (timestamps []int64, buckets []*Bucket) {
	maxBuckets := int(
		math.Ceil(
			float64(f.endTime-f.startTime) / float64(f.downsample),
		),
	)
	if maxBuckets >= math.MaxInt32 {
		maxBuckets = f.maxBuckets
	}

	timestamps = make([]int64, maxBuckets)
	for i := range timestamps {
		timestamps[i] = f.startTime + f.downsample*int64(i)
	}
	buckets = make([]*Bucket, 0, len(f.buckets))
	for _, b := range f.buckets {
		values := make([]int64, maxBuckets)
		copy(values, b.Values)
		buckets = append(buckets, &Bucket{
			Key:    b.Key,
			Values: values,
		})
	}
	return
}

func (f *FrequencyMap) Put(intervalInput interval.Interval, data map[string][]any) {
	type FieldValues struct {
		Field  string
		Values []any
	}
	keys := make(map[string]struct{})
	resultTS := temporalset.EmptyWithCapacity(len(data))
	for field, values := range data {
		resultTS.AddSet(&FieldValues{
			Field:  field,
			Values: values,
		}, temporalset.NewIntervalSet(intervalInput))
	}
	resultTS.EachInterval(func(ival interval.Interval, values []any) {
		data := make(map[string][]any, len(values))
		for _, v := range values {
			fv := v.(*FieldValues)
			data[fv.Field] = fv.Values
		}

		// now get every permutation
		options := make([]int, len(f.groupBy))
		for i, field := range f.groupBy {
			options[i] = len(data[field])
		}

		permute(options, func(index []int) {
			indexValues := make([]any, len(index))
			indexKeys := make(map[string]any, len(index))
			for i, j := range index {
				field := f.groupBy[i]
				if j < len(data[field]) {
					value := data[field][j]
					indexKeys[field] = value
					indexValues[i] = value
				}
			}
			// If grouping counts by field values and
			// the value key is empty, do not increment
			if len(f.groupBy) > 0 && len(indexKeys) == 0 {
				return
			}
			key := ToValues(indexValues...).String()
			if f.persistCounts {
				f.increment(key, indexKeys, ival)
			} else if _, ok := keys[key]; !ok {
				f.increment(key, indexKeys, ival)
				keys[key] = struct{}{}
			}
		})
	})
	return
}

func (f *FrequencyMap) increment(key string, values map[string]any, ival interval.Interval) {
	ival = ival.Intersect(interval.RightOpen(uint64(f.startTime), uint64(f.endTime)))
	if ival.IsEmpty() {
		return
	}
	lower, upper := ival.RightOpenBoundsInt64()

	var maxIndex int
	if f.persistCounts {
		maxIndex = int((upper - f.startTime) / f.downsample)
	} else {
		maxIndex = int((lower - f.startTime) / f.downsample)
	}

	for {
		f.mu.RLock()
		bucket, ok := f.buckets[key]
		if ok {
			if maxIndex < len(bucket.Values) {
				for i := int((lower - f.startTime) / f.downsample); i <= maxIndex; i++ {
					atomic.AddInt64(&bucket.Values[i], 1)
				}
				f.mu.RUnlock()
				return
			}
		}
		f.mu.RUnlock()

		f.mu.Lock()
		if maxIndex >= f.maxBuckets {
			f.maxBuckets = maxIndex + 1
		}

		bucket, ok = f.buckets[key]
		if ok {
			if maxIndex >= len(bucket.Values) {
				values := bucket.Values
				bucket.Values = make([]int64, f.maxBuckets)
				copy(bucket.Values, values)
			}
		} else {
			f.buckets[key] = &Bucket{
				Key:    values,
				Values: make([]int64, f.maxBuckets),
			}
		}
		f.mu.Unlock()
	}
}

func permute(input []int, f func(output []int)) {
	index := make([]int, len(input))
	f(index)

	for i := 0; i < len(input); {
		if index[i]+1 < input[i] {
			index[i]++
			for i > 0 {
				i--
				index[i] = 0
			}
			f(index)
		} else {
			i++
		}
	}
}
