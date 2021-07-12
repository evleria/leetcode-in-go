package _459_repeated_substring_pattern

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRepeatedSubstringPattern(t *testing.T) {
	testCases := []struct {
		gotNums string
		want    bool
	}{
		{
			gotNums: "abab",
			want:    true,
		},
		{
			gotNums: "aba",
			want:    false,
		},
	}

	for _, testCase := range testCases {
		actual := repeatedSubstringPattern(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
