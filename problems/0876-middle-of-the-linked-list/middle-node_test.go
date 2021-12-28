package _876_middle_of_the_linked_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMiddleNode(t *testing.T) {
	testCases := []struct {
		gotHead []int
		want    []int
	}{
		{
			gotHead: []int{1, 2, 3, 4, 5},
			want:    []int{3, 4, 5},
		},
		{
			gotHead: []int{1, 2, 3, 4, 5, 6},
			want:    []int{4, 5, 6},
		},
	}

	for _, testCase := range testCases {
		actual := middleNode(LinkedListFromSlice(testCase.gotHead)).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
