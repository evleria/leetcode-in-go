package _561_array_partition_I

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestArrayPairSum(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{1, 4, 3, 2},
			want: 4,
		},
		{
			got:  []int{6, 2, 6, 5, 1, 2},
			want: 9,
		},
	}

	for _, testCase := range testCases {
		actual := arrayPairSum(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
