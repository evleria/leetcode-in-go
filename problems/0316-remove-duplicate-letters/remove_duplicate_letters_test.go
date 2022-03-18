package _316_remove_duplicate_letters

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRemoveDuplicateLetters(t *testing.T) {
	testCases := []struct {
		gotNums string
		want    string
	}{
		{
			gotNums: "bcabc",
			want:    "abc",
		},
		{
			gotNums: "cbacdcbc",
			want:    "acdb",
		},
	}

	for _, testCase := range testCases {
		actual := removeDuplicateLetters(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
