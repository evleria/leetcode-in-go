package _143_longest_common_subsequence

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLongestCommonSubsequence(t *testing.T) {
	testCases := []struct {
		got1 string
		got2 string
		want int
	}{
		{
			got1: "abcde",
			got2: "ace",
			want: 3,
		},
		{
			got1: "abc",
			got2: "abc",
			want: 3,
		},
		{
			got1: "abc",
			got2: "def",
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := longestCommonSubsequence(testCase.got1, testCase.got2)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
