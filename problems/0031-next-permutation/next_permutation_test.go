package _031_next_permutation

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNextPermutation(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    []int
	}{
		{
			gotNums: []int{1, 2, 3},
			want:    []int{1, 3, 2},
		},
		{
			gotNums: []int{3, 2, 1},
			want:    []int{1, 2, 3},
		},
		{
			gotNums: []int{1, 1, 5},
			want:    []int{1, 5, 1},
		},
		{
			gotNums: []int{1},
			want:    []int{1},
		},
	}

	for _, testCase := range testCases {
		nextPermutation(testCase.gotNums)

		assert.Check(t, is.DeepEqual(testCase.gotNums, testCase.want))
	}
}
