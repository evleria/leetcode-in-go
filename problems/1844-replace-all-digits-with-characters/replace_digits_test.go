package _844_replace_all_digits_with_characters

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestReplaceDigits(t *testing.T) {
	testCases := []struct {
		got  string
		want string
	}{
		{
			got:  "a1c1e1",
			want: "abcdef",
		},
		{
			got:  "a1b2c3d4e",
			want: "abbdcfdhe",
		},
	}

	for _, testCase := range testCases {
		actual := replaceDigits(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
