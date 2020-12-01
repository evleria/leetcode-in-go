package _617_merge_two_binary_trees

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMergeTrees(t *testing.T) {
	testCases := []struct {
		gotTree1 []int
		gotTree2 []int
		want     []int
	}{
		{
			gotTree1: []int{1, 3, 2, 5},
			gotTree2: []int{2, 1, 3, NULL, 4, NULL, 7},
			want:     []int{3, 4, 5, 5, 4, NULL, 7},
		},
	}

	for _, testCase := range testCases {
		actual := mergeTrees(FromSlice(testCase.gotTree1), FromSlice(testCase.gotTree2)).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
