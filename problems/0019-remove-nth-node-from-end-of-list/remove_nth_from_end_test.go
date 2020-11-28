package _019_remove_nth_node_from_end_of_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRemoveNthFromEnd(t *testing.T) {
	testCases := []struct {
		gotList []int
		gotN    int
		want    []int
	}{
		{
			gotList: []int{1, 2, 3, 4, 5},
			gotN:    2,
			want:    []int{1, 2, 3, 5},
		},
		{
			gotList: []int{1},
			gotN:    1,
			want:    []int{},
		},
	}

	for _, testCase := range testCases {
		actual := removeNthFromEnd(FromSlice(testCase.gotList), testCase.gotN).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
