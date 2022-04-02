package _680_valid_palindrome_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestValidPalindrome(t *testing.T) {
	testCases := []struct {
		got  string
		want bool
	}{
		{
			got:  "aba",
			want: true,
		},
		{
			got:  "abca",
			want: true,
		},
		{
			got:  "abc",
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := validPalindrome(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
