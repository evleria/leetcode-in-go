package _108_find_first_palindromic_string_in_the_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFirstPalindrome(t *testing.T) {
	testCases := []struct {
		got  []string
		want string
	}{
		{
			got:  []string{"abc", "car", "ada", "racecar", "cool"},
			want: "ada",
		},
		{
			got:  []string{"notapalindrome", "racecar"},
			want: "racecar",
		},
		{
			got:  []string{"def", "ghi"},
			want: "",
		},
	}

	for _, testCase := range testCases {
		actual := firstPalindrome(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
