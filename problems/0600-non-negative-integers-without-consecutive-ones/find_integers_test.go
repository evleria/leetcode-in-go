package _600_non_negative_integers_without_consecutive_ones

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindIntegers(t *testing.T) {
	testCases := []struct {
		got  int
		want int
	}{
		{
			got:  5,
			want: 5,
		},
		{
			got:  1,
			want: 2,
		},
		{
			got:  2,
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := findIntegers(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
