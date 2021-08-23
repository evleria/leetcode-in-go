package _653_two_sum_IV_input_is_a_bst

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindTarget(t *testing.T) {
	testCases := []struct {
		gotRoot []int
		gotK    int
		want    bool
	}{
		{
			gotRoot: []int{45, 3, 6, 2, 4, NULL, 7},
			gotK:    9,
			want:    true,
		},
		{
			gotRoot: []int{5, 3, 6, 2, 4, NULL, 7},
			gotK:    28,
			want:    false,
		},
		{
			gotRoot: []int{2, 1, 3},
			gotK:    4,
			want:    true,
		},
	}

	for _, testCase := range testCases {
		actual := findTarget(BinaryTreeFromSlice(testCase.gotRoot), testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
