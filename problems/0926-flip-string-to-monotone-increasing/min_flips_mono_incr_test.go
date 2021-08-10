package _926_flip_string_to_monotone_increasing

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinFlipsMonoIncr(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "00110",
			want: 1,
		},
		{
			got:  "010110",
			want: 2,
		},
		{
			got:  "00011000",
			want: 2,
		},
	}

	for _, testCase := range testCases {
		actual := minFlipsMonoIncr(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
