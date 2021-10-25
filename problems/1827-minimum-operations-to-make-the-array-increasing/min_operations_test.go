package _827_minimum_operations_to_make_the_array_increasing

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinOperations(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{1, 1, 1},
			want: 3,
		},
		{
			got:  []int{1, 5, 2, 4, 1},
			want: 14,
		},
		{
			got:  []int{8},
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := minOperations(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
