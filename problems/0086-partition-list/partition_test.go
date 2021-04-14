package _086_partition_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestPartition(t *testing.T) {
	testCases := []struct {
		gotList []int
		gotX    int
		want    []int
	}{
		{
			gotList: []int{1, 4, 3, 2, 5, 2},
			gotX:    3,
			want:    []int{1, 2, 2, 4, 3, 5},
		},
		{
			gotList: []int{2, 1},
			gotX:    2,
			want:    []int{1, 2},
		},
	}

	for _, testCase := range testCases {
		actual := partition(FromSlice(testCase.gotList), testCase.gotX)

		assert.Check(t, is.DeepEqual(actual.ToSlice(), testCase.want))
	}
}
