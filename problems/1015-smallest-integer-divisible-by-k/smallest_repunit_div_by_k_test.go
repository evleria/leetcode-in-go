package _015_smallest_integer_divisible_by_k

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSmallestRepunitDivByK(t *testing.T) {
	testCases := []struct {
		got  int
		want int
	}{
		{
			got:  1,
			want: 1,
		},
		{
			got:  2,
			want: -1,
		},
		{
			got:  3,
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := smallestRepunitDivByK(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
