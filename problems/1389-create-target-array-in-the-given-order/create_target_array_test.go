package _389_create_target_array_in_the_given_order

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCreateTargetArray(t *testing.T) {
	testCases := []struct {
		gotNums []int
		gotIdx  []int
		want    []int
	}{
		{
			gotNums: []int{0, 1, 2, 3, 4},
			gotIdx:  []int{0, 1, 2, 2, 1},
			want:    []int{0, 4, 1, 3, 2},
		},
		{
			gotNums: []int{1, 2, 3, 4, 0},
			gotIdx:  []int{0, 1, 2, 3, 0},
			want:    []int{0, 1, 2, 3, 4},
		},
		{
			gotNums: []int{1},
			gotIdx:  []int{0},
			want:    []int{1},
		},
	}

	for _, testCase := range testCases {
		actual := createTargetArray(testCase.gotNums, testCase.gotIdx)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
