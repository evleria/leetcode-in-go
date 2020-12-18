package _237_delete_node_in_a_linked_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestDeleteNode(t *testing.T) {
	testCases := []struct {
		gotList      []int
		gotNodeIndex int
		want         []int
	}{
		{
			gotList:      []int{1, 2, 3, 4},
			gotNodeIndex: 2,
			want:         []int{1, 2, 4},
		},
	}

	for _, testCase := range testCases {
		list := FromSlice(testCase.gotList)
		node, _ := list.NodeAt(testCase.gotNodeIndex)

		deleteNode(node)

		assert.Check(t, is.DeepEqual(list.ToSlice(), testCase.want))
	}
}
