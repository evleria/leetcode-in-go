package _409_longest_palindrome

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "abccccdd",
			want: 7,
		},
		{
			got:  "a",
			want: 1,
		},
		{
			got:  "bb",
			want: 2,
		},
	}

	for _, testCase := range testCases {
		actual := longestPalindrome(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
