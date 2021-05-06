package _206_reverse_linked_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestReverseList(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{1, 2, 3, 4},
			want: []int{4, 3, 2, 1},
		},
	}

	for _, testCase := range testCases {
		actual := reverseList(LinkedListFromSlice(testCase.got)).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
