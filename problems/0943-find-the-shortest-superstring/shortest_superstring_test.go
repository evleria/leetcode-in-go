package _943_find_the_shortest_superstring

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"

	"testing"
)

func TestShortestSuperstring(t *testing.T) {
	testCases := []struct {
		got  []string
		want string
	}{
		{
			got:  []string{"abc", "cde", "de"},
			want: "abcde",
		},
		{
			got:  []string{"alex", "loves", "leetcode"},
			want: "leetcodelovesalex",
		},
		{
			got:  []string{"catg", "ctaagt", "gcta", "ttca", "atgcatc"},
			want: "gctaagttcatgcatc",
		},
	}

	for _, testCase := range testCases {
		actual := shortestSuperstring(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
