package _022_generate_parentheses

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		got  int
		want []string
	}{
		{
			got:  1,
			want: []string{"()"},
		},
		{
			got:  3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}

	for _, testCase := range testCases {
		actual := generateParenthesis(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
