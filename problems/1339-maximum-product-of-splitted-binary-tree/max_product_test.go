package _339_maximum_product_of_splitted_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxProduct(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{1, 2, 3, 4, 5, 6},
			want: 110,
		},
		{
			got:  []int{1, NULL, 2, 3, 4, NULL, NULL, 5, 6},
			want: 90,
		},
		{
			got:  []int{2, 3, 9, 10, 7, 8, 6, 5, 4, 11, 1},
			want: 1025,
		},
		{
			got:  []int{1, 1},
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := maxProduct(BinaryTreeFromSlice(testCase.got))

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
