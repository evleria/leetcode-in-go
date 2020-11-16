package _101_symmetric_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsSymmetric(t *testing.T) {
	testCases := []struct {
		got  []int
		want bool
	}{
		{
			got:  []int{},
			want: true,
		},
		{
			got:  []int{1, 2, 2, 3, 4, 4, 3},
			want: true,
		},
		{
			got:  []int{1, 2, 2, NULL, 3, NULL, 3},
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := isSymmetric(FromSlice(testCase.got))

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
