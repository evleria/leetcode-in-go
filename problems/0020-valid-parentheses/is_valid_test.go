package _020_valid_parentheses

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		got  string
		want bool
	}{
		{
			got:  "[]",
			want: true,
		},
		{
			got:  "{",
			want: false,
		},
		{
			got:  "{}[]()",
			want: true,
		},
		{
			got:  "{[()()]}",
			want: true,
		},
		{
			got:  "{[]})",
			want: false,
		},
		{
			got:  "[(])",
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := isValid(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
