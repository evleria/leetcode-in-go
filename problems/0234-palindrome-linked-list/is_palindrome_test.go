package _234_palindrome_linked_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		got  []int
		want bool
	}{
		{
			got:  []int{1},
			want: true,
		},
		{
			got:  []int{1, 2},
			want: false,
		},
		{
			got:  []int{1, 2, 2, 1},
			want: true,
		},
		{
			got:  []int{1, 2, 3, 2, 1},
			want: true,
		},
	}

	for _, testCase := range testCases {
		actual := isPalindrome(LinkedListFromSlice(testCase.got))

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
