package _328_odd_even_linked_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestOddEvenList(t *testing.T) {
	testCases := []struct {
		gotHead []int
		want    []int
	}{
		{
			gotHead: []int{1, 2, 3, 4, 5},
			want:    []int{1, 3, 5, 2, 4},
		},
		{
			gotHead: []int{2, 1, 3, 5, 6, 4, 7},
			want:    []int{2, 3, 6, 7, 1, 5, 4},
		},
	}

	for _, testCase := range testCases {
		actual := oddEvenList(LinkedListFromSlice(testCase.gotHead)).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
