package _915_partition_array_into_disjoint_intervals

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestPartitionDisjoint(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{32, 57, 24, 19, 0, 24, 49, 67, 87, 87},
			want:    7,
		},
		{
			gotNums: []int{5, 0, 3, 8, 6},
			want:    3,
		},
		{
			gotNums: []int{1, 1, 1, 0, 6, 12},
			want:    4,
		},
	}

	for _, testCase := range testCases {
		actual := partitionDisjoint(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
