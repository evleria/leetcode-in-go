package _328_break_a_palindrome

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestBreakPalindrome(t *testing.T) {
	testCases := []struct {
		gotS string
		want string
	}{
		{
			gotS: "abccba",
			want: "aaccba",
		},
		{
			gotS: "a",
			want: "",
		},
		{
			gotS: "aa",
			want: "ab",
		},
		{
			gotS: "aba",
			want: "abb",
		},
	}

	for _, testCase := range testCases {
		actual := breakPalindrome(testCase.gotS)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
