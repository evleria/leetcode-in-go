package _009_palindrome_number

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		got  int
		want bool
	}{
		{
			got:  121,
			want: true,
		},
		{
			got:  1221,
			want: true,
		},
		{
			got:  -121,
			want: false,
		},
		{
			got:  10,
			want: false,
		},
		{
			got:  -101,
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := isPalindrome(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
