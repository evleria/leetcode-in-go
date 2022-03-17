package _856_score_of_parentheses

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestScoreOfParentheses(t *testing.T) {
	testCases := []struct {
		gotNums string
		want    int
	}{
		{
			gotNums: "()",
			want:    1,
		},
		{
			gotNums: "(())",
			want:    2,
		},
		{
			gotNums: "()()",
			want:    2,
		},
	}

	for _, testCase := range testCases {
		actual := scoreOfParentheses(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
