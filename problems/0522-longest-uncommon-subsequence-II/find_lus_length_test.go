package _522_longest_uncommon_subsequence_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindLUSLength(t *testing.T) {
	testCases := []struct {
		got  []string
		want int
	}{
		{
			got:  []string{"aba", "cdc", "eae"},
			want: 3,
		},
		{
			got:  []string{"aaa", "aaa", "aa"},
			want: -1,
		},
	}

	for _, testCase := range testCases {
		actual := findLUSlength(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
