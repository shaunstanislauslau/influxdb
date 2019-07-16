package tsm1

import (
	"math"
	"testing"

	"github.com/influxdata/influxdb/tsdb/cursors"
)

func TestCombineFloat(t *testing.T) {

	key := []byte("foo")

	floats := func(min, max, count int64) []byte {
		delta := (max - min) / count
		a := cursors.NewFloatArrayLen(0)

		for ts := min; ts <= max && a.Len() < int(count); ts += delta {
			a.Timestamps = append(a.Timestamps, ts)
			a.Values = append(a.Values, 1.0)
		}
		if a.Timestamps[len(a.Timestamps)-1] < max {
			a.Timestamps[len(a.Timestamps)-1] = max
		}

		buf, err := EncodeFloatArrayBlock(a, nil)
		if err != nil {
			panic(err)
		}
		return buf
	}

	blk := func(min, max, count int64) *block {
		return &block{
			key:        key,
			minTime:    min,
			maxTime:    max,
			typ:        BlockFloat64,
			b:          floats(min, max, count),
			tombstones: nil,
			readMin:    math.MaxInt64,
			readMax:    math.MinInt64,
		}
	}

	iter := tsmBatchKeyIterator{
		size: MaxPointsPerBlock,
		key:  key,
		typ:  BlockFloat64,
		blocks: []*block{
			blk(10000, 12000, 1000),
			blk(10500, 11500, 1000),
		},
		mergedFloatValues: cursors.NewFloatArrayLen(0),
	}

	iter.mergeFloat()
	iter.mergeFloat()
}
