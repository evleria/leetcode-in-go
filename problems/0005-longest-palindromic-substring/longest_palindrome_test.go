package _005_longest_palindromic_substring

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		got  string
		want string
	}{
		{
			got:  "babad",
			want: "bab",
		},
		{
			got:  "cbbd",
			want: "bb",
		},
	}
	for _, testCase := range testCases {
		actual := longestPalindrome(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
