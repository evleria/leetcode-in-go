package _979_find_greatest_common_divisor_of_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindGCD(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{2, 5, 6, 9, 10},
			want: 2,
		},
		{
			got:  []int{7, 5, 6, 8, 3},
			want: 1,
		},
		{
			got:  []int{3, 3},
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := findGCD(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
