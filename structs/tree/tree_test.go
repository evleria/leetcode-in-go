package tree

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFromToSlice(t *testing.T) {
	testCases := [][]int{
		{},
		{0, 1, 4, 8},
		{2, 1, 3, NULL, 4, NULL, 7},
		{3, 4, 5, 5, 4, NULL, 7},
	}

	for _, testCase := range testCases {
		tree := FromSlice(testCase)
		slice := tree.ToSlice()

		assert.Check(t, is.DeepEqual(slice, testCase))
	}
}
