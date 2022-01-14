package _008_string_to_integer

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMyAtoi(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "42",
			want: 42,
		},
		{
			got:  "   -42",
			want: -42,
		},
		{
			got:  "4193 with words",
			want: 4193,
		},
	}

	for _, testCase := range testCases {
		actual := myAtoi(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
