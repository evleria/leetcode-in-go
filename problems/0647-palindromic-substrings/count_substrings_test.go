package _647_palindromic_substrings

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCountSubstrings(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "abc",
			want: 3,
		},
		{
			got:  "aaa",
			want: 6,
		},
	}

	for _, testCase := range testCases {
		actual := countSubstrings(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
