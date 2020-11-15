package _100_is_same_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsSameTree(t *testing.T) {
	testCases := []struct {
		got1 []int
		got2 []int
		want bool
	}{
		{
			got1: []int{1, 2, 3},
			got2: []int{1, 2, 3},
			want: true,
		},
		{
			got1: []int{1, 2},
			got2: []int{1, NULL, 2},
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := isSameTree(FromSlice(testCase.got1), FromSlice(testCase.got2))

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
