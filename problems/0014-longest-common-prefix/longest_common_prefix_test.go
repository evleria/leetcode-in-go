package _014_longest_common_prefix

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		got  []string
		want string
	}{
		{
			got:  []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			got:  []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			got:  []string{"racecar", "race", "racer"},
			want: "race",
		},
		{
			got:  []string{},
			want: "",
		},
	}

	for _, testCase := range testCases {
		actual := longestCommonPrefix(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
