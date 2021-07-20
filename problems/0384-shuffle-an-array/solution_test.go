package _384_shuffle_an_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSolution(t *testing.T) {
	solution := Constructor([]int{1, 2, 3})

	assert.Check(t, is.DeepEqual(solution.Shuffle(), []int{1, 3, 2}))
	assert.Check(t, is.DeepEqual(solution.Reset(), []int{1, 2, 3}))
	assert.Check(t, is.DeepEqual(solution.Shuffle(), []int{3, 1, 2}))
}
