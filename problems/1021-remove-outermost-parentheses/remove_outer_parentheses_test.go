package _021_remove_outermost_parentheses

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"

	"testing"
)

func TestRemoveOuterParentheses(t *testing.T) {
	testCases := []struct {
		got  string
		want string
	}{
		{
			got:  "(()())(())",
			want: "()()()",
		},
		{
			got:  "(()())(())(()(()))",
			want: "()()()()(())",
		},
		{
			got:  "()()",
			want: "",
		},
	}

	for _, testCase := range testCases {
		actual := removeOuterParentheses(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
