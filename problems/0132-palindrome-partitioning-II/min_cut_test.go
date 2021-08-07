package _132_palindrome_partitioning_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinCut(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "aaabaa",
			want: 1,
		},
		{
			got:  "aab",
			want: 1,
		},
		{
			got:  "a",
			want: 0,
		},
		{
			got:  "ab",
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := minCut(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
