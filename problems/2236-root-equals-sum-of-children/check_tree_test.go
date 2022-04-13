package _236_root_equals_sum_of_children

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCheckTree(t *testing.T) {
	testCases := []struct {
		got  []int
		want bool
	}{
		{
			got:  []int{10, 4, 6},
			want: true,
		},
		{
			got:  []int{5, 3, 1},
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := checkTree(BinaryTreeFromSlice(testCase.got))

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
