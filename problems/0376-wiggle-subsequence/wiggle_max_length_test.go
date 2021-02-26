package _376_wiggle_subsequence

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestWiggleMaxLength(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{},
			want: 0,
		},
		{
			got:  []int{1},
			want: 1,
		},
		{
			got:  []int{1, 1},
			want: 1,
		},
		{
			got:  []int{1, 7, 4, 9, 2, 5},
			want: 6,
		},
		{
			got:  []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8},
			want: 7,
		},
	}

	for _, testCase := range testCases {
		actual := wiggleMaxLength(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
