package _978_longest_turbulent_subarray

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxTurbulenceSize(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{9, 4, 2, 10, 7, 8, 8, 1, 9},
			want: 5,
		},
		{
			got:  []int{4, 8, 12, 16},
			want: 2,
		},
		{
			got:  []int{100},
			want: 1,
		},
		{
			got:  []int{1, 2, 3},
			want: 2,
		},
	}

	for _, testCase := range testCases {
		actual := maxTurbulenceSize(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
