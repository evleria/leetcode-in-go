package _303_range_sum_query_immutable

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumArray(t *testing.T) {
	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})
	assert.Check(t, is.Equal(numArray.SumRange(0, 2), 1))
	assert.Check(t, is.Equal(numArray.SumRange(2, 5), -1))
	assert.Check(t, is.Equal(numArray.SumRange(0, 5), -3))
}
