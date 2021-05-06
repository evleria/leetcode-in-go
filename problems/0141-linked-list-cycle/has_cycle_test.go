package _141_linked_list_cycle

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestHasCycle(t *testing.T) {
	testCases := []struct {
		gotList    []int
		gotCycleTo int
		want       bool
	}{
		{
			gotList:    []int{3, 2, 0, -4},
			gotCycleTo: 1,
			want:       true,
		},
		{
			gotList:    []int{1, 2},
			gotCycleTo: 0,
			want:       true,
		},
		{
			gotList:    []int{1, 2},
			gotCycleTo: -1,
			want:       false,
		},
	}

	for _, testCase := range testCases {
		list := LinkedListFromSlice(testCase.gotList)

		if testCase.gotCycleTo != -1 {
			lastNode, _ := list.NodeAt(len(testCase.gotList) - 1)
			toNode, _ := list.NodeAt(testCase.gotCycleTo)
			lastNode.Next = toNode
		}

		actual := hasCycle(list)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
