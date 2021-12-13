package _446_consecutive_characters

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxPower(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "leetcode",
			want: 2,
		},
		{
			got:  "abbcccddddeeeeedcba",
			want: 5,
		},
		{
			got:  "triplepillooooow",
			want: 5,
		},
		{
			got:  "hooraaaaaaaaaaay",
			want: 11,
		},
		{
			got:  "tourist",
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := maxPower(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
