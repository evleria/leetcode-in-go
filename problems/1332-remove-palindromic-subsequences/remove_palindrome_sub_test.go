package _332_remove_palindromic_subsequences

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "",
			want: 0,
		},
		{
			got:  "ababa",
			want: 1,
		},
		{
			got:  "abb",
			want: 2,
		},
	}

	for _, testCase := range testCases {
		actual := removePalindromeSub(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
