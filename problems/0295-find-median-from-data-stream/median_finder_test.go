package _295_find_median_from_data_stream

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMedianFinder(t *testing.T) {
	finder := Constructor()
	finder.AddNum(1)
	finder.AddNum(2)
	assert.Check(t, is.Equal(finder.FindMedian(), 1.5))
	finder.AddNum(3)
	assert.Check(t, is.Equal(finder.FindMedian(), 2.0))
}
