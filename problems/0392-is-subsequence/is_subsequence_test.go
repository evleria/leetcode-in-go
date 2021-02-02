package _392_is_subsequence

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsSubsequence(t *testing.T) {
	testCases := []struct {
		got1 string
		got2 string
		want bool
	}{
		{
			got1: "abc",
			got2: "ahbgdc",
			want: true,
		},
		{
			got1: "axc",
			got2: "ahbgdc",
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := isSubsequence(testCase.got1, testCase.got2)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
