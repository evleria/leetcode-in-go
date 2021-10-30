package _044_longest_duplicate_substring

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestLongestDubSubstring(t *testing.T) {
	testCases := []struct {
		got  string
		want string
	}{
		{
			got:  "banana",
			want: "ana",
		},
		{
			got:  "abcd",
			want: "",
		},
		{
			got:  "aa",
			want: "a",
		},
	}

	for _, testCase := range testCases {
		actual := longestDupSubstring(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
