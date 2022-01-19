package _142_linked_list_cycle_II

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestDetectCycle(t *testing.T) {
	testCases := []struct {
		gotList    []int
		gotCycleTo int
		want       bool
	}{
		{
			gotList:    []int{3, 2, 0, -4},
			gotCycleTo: 1,
		},
		{
			gotList:    []int{1, 2},
			gotCycleTo: 0,
		},
		{
			gotList:    []int{1, 2},
			gotCycleTo: -1,
		},
	}

	for _, testCase := range testCases {
		list := LinkedListFromSlice(testCase.gotList)

		if testCase.gotCycleTo != -1 {
			lastNode, _ := list.NodeAt(len(testCase.gotList) - 1)
			toNode, _ := list.NodeAt(testCase.gotCycleTo)
			lastNode.Next = toNode
		}

		actual := detectCycle(list)

		if testCase.gotCycleTo != -1 {
			want, _ := list.NodeAt(testCase.gotCycleTo)
			assert.Check(t, is.Equal(actual, want))
		} else {
			assert.Check(t, is.Nil(actual))
		}
	}
}
