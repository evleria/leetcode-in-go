package _430_flatten_a_multilevel_doubly_linked_list

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestFlatten(t *testing.T) {
	testCases := []struct {
		got  *Node
		want []int
	}{
		{
			got: withAttached(fromSlice([]int{1, 2, 3, 4, 5, 6}), 2,
				withAttached(fromSlice([]int{7, 8, 9, 10}), 1,
					fromSlice([]int{11, 12}))),

			want: []int{1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6},
		},
	}

	for _, testCase := range testCases {
		actual := flatten(testCase.got)

		assert.Check(t, is.DeepEqual(toFlatSlice(actual), testCase.want))
	}
}

func fromSlice(nums []int) *Node {
	dummyHead := new(Node)

	cur := dummyHead
	for _, n := range nums {
		cur.Next = &Node{Val: n, Prev: cur}
		cur = cur.Next
	}

	return dummyHead.Next
}

func withAttached(node *Node, index int, child *Node) *Node {
	cur := node
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Child = child

	return node
}

func toFlatSlice(node *Node) []int {
	result := []int{}

	for node != nil {
		if node.Child != nil {
			panic("node has child")
		}

		result = append(result, node.Val)
		node = node.Next
	}

	return result
}
