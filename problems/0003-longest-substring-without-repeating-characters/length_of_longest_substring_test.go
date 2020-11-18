package _003_longest_substring_without_repeating_characters

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "abcabcbb",
			want: 3,
		},
		{
			got:  "bbbbb",
			want: 1,
		},
		{
			got:  "pwwkew",
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := lengthOfLongestSubstring(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
