package _305_all_elements_in_two_binary_search_trees

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGetAllElements(t *testing.T) {
	testCases := []struct {
		got1 []int
		got2 []int
		want []int
	}{
		{
			got1: []int{2, 1, 4},
			got2: []int{1, 0, 3},
			want: []int{0, 1, 1, 2, 3, 4},
		},
		{
			got1: []int{1, NULL, 8},
			got2: []int{8, 1},
			want: []int{1, 1, 8, 8},
		},
	}

	for _, testCase := range testCases {
		actual := getAllElements(BinaryTreeFromSlice(testCase.got1), BinaryTreeFromSlice(testCase.got2))

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
