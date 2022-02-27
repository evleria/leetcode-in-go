package _662_maximum_width_of_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

type NodeInfo struct {
	node *TreeNode
	idx  int
}

func widthOfBinaryTree(root *TreeNode) int {
	queue := []NodeInfo{{root, 0}}

	maxWidth := 0
	for len(queue) > 0 {
		min, max := queue[0].idx, 0
		for _, cur := range queue {
			queue = queue[1:]
			max = cur.idx
			if cur.node.Left != nil {
				queue = append(queue, NodeInfo{cur.node.Left, 2 * cur.idx})
			}

			if cur.node.Right != nil {
				queue = append(queue, NodeInfo{cur.node.Right, 2*cur.idx + 1})
			}
		}

		if width := max - min + 1; width > maxWidth {
			maxWidth = width
		}
	}

	return maxWidth
}
