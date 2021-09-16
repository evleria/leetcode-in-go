package _125_valid_palindrome

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		gotNums string
		want    bool
	}{
		{
			gotNums: "A man, a plan, a canal: Panama",
			want:    true,
		},
		{
			gotNums: "race a car",
			want:    false,
		},
		{
			gotNums: "0P",
			want:    false,
		},
	}

	for _, testCase := range testCases {
		actual := isPalindrome(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
