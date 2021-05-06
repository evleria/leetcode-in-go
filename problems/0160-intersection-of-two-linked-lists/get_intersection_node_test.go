package _160_intersection_of_two_linked_lists

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGetIntersectionNode(t *testing.T) {
	testCases := []struct {
		gotListA []int
		gotListB []int
		gotSkipA int
		gotSkipB int
		want     int
	}{
		{
			gotListA: []int{4, 1, 8, 4, 5},
			gotListB: []int{5, 6, 1, 8, 4, 5},
			gotSkipA: 2,
			gotSkipB: 3,
			want:     8,
		},
		{
			gotListA: []int{2, 6, 4},
			gotListB: []int{1, 5},
			gotSkipA: 0,
			gotSkipB: 0,
			want:     0,
		},
	}

	for _, testCase := range testCases {
		listA := LinkedListFromSlice(testCase.gotListA)
		listB := LinkedListFromSlice(testCase.gotListB)
		if testCase.gotSkipA != 0 && testCase.gotSkipB != 0 {
			targetNode, _ := listA.NodeAt(testCase.gotSkipA)
			sourceNode, _ := listB.NodeAt(testCase.gotSkipB - 1)
			sourceNode.Next = targetNode

			actual := getIntersectionNode(listA, listB)

			assert.Check(t, is.Equal(actual.Val, testCase.want))
		} else {
			actual := getIntersectionNode(listA, listB)

			assert.Check(t, is.Nil(actual))
		}
	}
}
