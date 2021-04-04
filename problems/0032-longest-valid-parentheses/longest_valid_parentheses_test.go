package _032_longest_valid_parentheses

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "",
			want: 0,
		},
		{
			got:  "(()",
			want: 2,
		},
		{
			got:  ")()())",
			want: 4,
		},
	}

	for _, testCase := range testCases {
		actual := longestValidParentheses(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
