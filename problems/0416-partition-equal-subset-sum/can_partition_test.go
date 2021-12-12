package _416_partition_equal_subset_sum

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCanPartition(t *testing.T) {
	testCases := []struct {
		got  []int
		want bool
	}{
		{
			got:  []int{1, 5, 11, 5},
			want: true,
		},
		{
			got:  []int{1, 2, 3, 5},
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := canPartition(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
