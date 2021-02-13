package _150_evaluate_reverse_polish_notation

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestEvalRPN(t *testing.T) {
	testCases := []struct {
		got  []string
		want int
	}{
		{
			got:  []string{"2", "1", "+", "3", "*"},
			want: 9,
		},
		{
			got:  []string{"4", "13", "5", "/", "+"},
			want: 6,
		},
		{
			got:  []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want: 22,
		},
	}

	for _, testCase := range testCases {
		actual := evalRPN(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
