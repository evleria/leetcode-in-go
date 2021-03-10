package _012_integer_to_roman

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIntToRoman(t *testing.T) {
	testCases := []struct {
		got  int
		want string
	}{
		{
			got:  3,
			want: "III",
		},
		{
			got:  4,
			want: "IV",
		},
		{
			got:  1994,
			want: "MCMXCIV",
		},
	}

	for _, testCase := range testCases {
		actual := intToRoman(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
