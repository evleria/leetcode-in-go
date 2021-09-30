package _698_partition_to_k_equal_sum_subsets

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCanPartitionKSubsets(t *testing.T) {
	testCases := []struct {
		got  []int
		gotK int
		want bool
	}{
		{
			got:  []int{4, 3, 2, 3, 5, 2, 1},
			gotK: 4,
			want: true,
		},
		{
			got:  []int{1, 2, 3, 4},
			gotK: 3,
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := canPartitionKSubsets(testCase.got, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
