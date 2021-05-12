package _304_range_sum_query_2d_immutable

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumMatrix(t *testing.T) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	numMatrix := Constructor(matrix)
	assert.Check(t, is.Equal(numMatrix.SumRegion(2, 1, 4, 3), 8))
	assert.Check(t, is.Equal(numMatrix.SumRegion(1, 1, 2, 2), 11))
	assert.Check(t, is.Equal(numMatrix.SumRegion(1, 2, 2, 4), 12))
}
