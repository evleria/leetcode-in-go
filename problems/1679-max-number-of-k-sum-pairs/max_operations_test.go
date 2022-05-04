package _679_max_number_of_k_sum_pairs

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxOperations(t *testing.T) {
	testCases := []struct {
		gotNums []int
		gotK    int
		want    int
	}{
		{
			gotNums: []int{1, 2, 3, 4},
			gotK:    5,
			want:    2,
		},
		{
			gotNums: []int{3, 1, 3, 4, 3},
			gotK:    6,
			want:    1,
		},
	}

	for _, testCase := range testCases {
		actual := maxOperations(testCase.gotNums, testCase.gotK)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
