package _800_maximum_ascending_subarray_sum

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxAscendingSum(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{10, 20, 30, 5, 10, 50},
			want: 65,
		},
		{
			got:  []int{10, 20, 30, 40, 50},
			want: 150,
		},
		{
			got:  []int{12, 17, 15, 13, 10, 11, 12},
			want: 33,
		},
	}

	for _, testCase := range testCases {
		actual := maxAscendingSum(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
