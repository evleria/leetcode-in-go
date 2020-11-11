package _203_remove_linked_list_elements

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRemoveElements(t *testing.T) {
	testCases := []struct {
		gotList []int
		gotVal  int
		want    []int
	}{
		{
			gotList: []int{1, 2, 6, 3, 4, 5, 6},
			gotVal:  6,
			want:    []int{1, 2, 3, 4, 5},
		},
	}

	for _, testCase := range testCases {
		actual := removeElements(FromSlice(testCase.gotList), testCase.gotVal).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
