package _696_count_binary_substrings

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCountBinarySubstrings(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "00110011",
			want: 6,
		},
		{
			got:  "10101",
			want: 4,
		},
	}

	for _, testCase := range testCases {
		actual := countBinarySubstrings(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
