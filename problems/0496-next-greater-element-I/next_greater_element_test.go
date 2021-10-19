package _496_next_greater_element_I

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNextGreaterElement(t *testing.T) {
	testCases := []struct {
		gotNums   []int
		gotTarget []int
		want      []int
	}{
		{
			gotNums:   []int{4, 1, 2},
			gotTarget: []int{1, 3, 4, 2},
			want:      []int{-1, 3, -1},
		},
		{
			gotNums:   []int{2, 4},
			gotTarget: []int{1, 2, 3, 4},
			want:      []int{3, -1},
		},
	}

	for _, testCase := range testCases {
		actual := nextGreaterElement(testCase.gotNums, testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
