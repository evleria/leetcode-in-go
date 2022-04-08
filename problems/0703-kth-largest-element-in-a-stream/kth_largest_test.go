package _703_kth_largest_element_in_a_stream

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestKthLargest(t *testing.T) {
	cache := Constructor(3, []int{4, 5, 8, 2})
	assert.Check(t, is.Equal(cache.Add(3), 4))
	assert.Check(t, is.Equal(cache.Add(5), 5))
	assert.Check(t, is.Equal(cache.Add(10), 5))
	assert.Check(t, is.Equal(cache.Add(9), 8))
	assert.Check(t, is.Equal(cache.Add(4), 8))
}
