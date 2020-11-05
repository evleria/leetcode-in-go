package _088_merge_sorted_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		gotNums1 []int
		gotM     int
		gotNums2 []int
		gotN     int
		want     []int
	}{
		{
			gotNums1: []int{0},
			gotM:     0,
			gotNums2: []int{1},
			gotN:     1,
			want:     []int{1},
		},
		{
			gotNums1: []int{1, 2, 3, 0, 0, 0},
			gotM:     3,
			gotNums2: []int{2, 5, 6},
			gotN:     3,
			want:     []int{1, 2, 2, 3, 5, 6},
		},
	}

	for _, testCase := range testCases {
		merge(testCase.gotNums1, testCase.gotM, testCase.gotNums2, testCase.gotN)
		assert.Check(t, is.DeepEqual(testCase.gotNums1, testCase.want))
	}
}
