package _006_count_number_of_pairs_with_absolute_difference_k

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCountKDifference(t *testing.T) {
	testCases := []struct {
		gotNums []int
		gotK    int
		want    int
	}{
		{
			gotNums: []int{1, 2, 2, 1},
			gotK:    1,
			want:    4,
		},
		{
			gotNums: []int{1, 3},
			gotK:    3,
			want:    0,
		},
		{
			gotNums: []int{3, 2, 1, 5, 4},
			gotK:    2,
			want:    3,
		},
	}

	for _, testCase := range testCases {
		actual := countKDifference(testCase.gotNums, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
