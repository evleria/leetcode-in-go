package _470_shuffle_the_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestShuffle(t *testing.T) {
	testCases := []struct {
		gotNums []int
		gotN    int
		want    []int
	}{
		{
			gotNums: []int{2, 5, 1, 3, 4, 7},
			gotN:    3,
			want:    []int{2, 3, 5, 4, 1, 7},
		},
		{
			gotNums: []int{1, 2, 3, 4, 4, 3, 2, 1},
			gotN:    4,
			want:    []int{1, 4, 2, 3, 3, 2, 4, 1},
		},
		{
			gotNums: []int{1, 1, 2, 2},
			gotN:    2,
			want:    []int{1, 2, 1, 2},
		},
	}

	for _, testCase := range testCases {
		actual := shuffle(testCase.gotNums, testCase.gotN)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
