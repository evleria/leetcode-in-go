package _525_contiguous_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindMaxLength(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{

		{
			got:  []int{0, 1},
			want: 2,
		},
		{
			got:  []int{0, 1, 0},
			want: 2,
		},
	}

	for _, testCase := range testCases {
		actual := findMaxLength(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
