package _249_minimum_remove_to_make_valid_parentheses

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinRemoveToMakeValid(t *testing.T) {
	testCases := []struct {
		got  string
		want string
	}{
		{
			got:  "lee(t(c)o)de)",
			want: "lee(t(c)o)de",
		},
		{
			got:  "a)b(c)d",
			want: "ab(c)d",
		},
		{
			got:  "))((",
			want: "",
		},
	}

	for _, testCase := range testCases {
		actual := minRemoveToMakeValid(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
