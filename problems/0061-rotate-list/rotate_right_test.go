package _061_rotate_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRotateRight(t *testing.T) {
	testCases := []struct {
		gotHead []int
		gotK    int
		want    []int
	}{
		{
			gotHead: []int{},
			gotK:    1,
			want:    []int{},
		},
		{
			gotHead: []int{1, 2, 3, 4, 5},
			gotK:    2,
			want:    []int{4, 5, 1, 2, 3},
		},
		{
			gotHead: []int{1, 2, 3},
			gotK:    4,
			want:    []int{3, 1, 2},
		},
		{
			gotHead: []int{1, 2, 3},
			gotK:    3,
			want:    []int{1, 2, 3},
		},
	}

	for _, testCase := range testCases {
		actual := rotateRight(FromSlice(testCase.gotHead), testCase.gotK).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
