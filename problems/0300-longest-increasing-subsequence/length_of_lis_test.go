package _300_longest_increasing_subsequence

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLengthOfLis(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4,
		},
		{
			got:  []int{0, 1, 0, 3, 2, 3},
			want: 4,
		},
		{
			got:  []int{7, 7, 7, 7, 7, 7, 7},
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := lengthOfLIS(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
