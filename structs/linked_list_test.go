package structs

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFromToSlice(t *testing.T) {
	testCases := [][]int{
		{},
		{0, 1, 4, 8},
	}

	for _, testCase := range testCases {
		list := FromSlice(testCase)
		slice := ToSlice(list)

		assert.Check(t, is.DeepEqual(slice, testCase))
	}
}
