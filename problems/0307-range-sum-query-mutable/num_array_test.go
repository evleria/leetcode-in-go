package _307_range_sum_query_mutable

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumArray(t *testing.T) {
	numArray := Constructor([]int{1, 3, 5})
	assert.Check(t, is.Equal(numArray.SumRange(0, 2), 9))
	numArray.Update(1, 2)
	assert.Check(t, is.Equal(numArray.SumRange(0, 2), 8))
}
