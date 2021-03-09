package _623_add_one_row_to_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestAddOneRow(t *testing.T) {
	testCases := []struct {
		gotRoot []int
		gotV    int
		gotD    int
		want    []int
	}{
		{
			gotRoot: []int{4, 2, 6, 3, 1, 5},
			gotV:    1,
			gotD:    2,
			want:    []int{4, 1, 1, 2, NULL, NULL, 6, 3, 1, 5},
		},
	}

	for _, testCase := range testCases {
		actual := addOneRow(FromSlice(testCase.gotRoot), testCase.gotV, testCase.gotD).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
