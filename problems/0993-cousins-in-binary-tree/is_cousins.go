package _993_cousins_in_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

type Node struct {
	Parent *TreeNode
	Child  *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	queue := []Node{{nil, root}}
	for len(queue) > 0 {
		var xParent, yParent *TreeNode
		for _, node := range queue {
			queue = queue[1:]

			if node.Child == nil {
				continue
			}

			if node.Child.Val == x {
				xParent = node.Parent
			} else if node.Child.Val == y {
				yParent = node.Parent
			}

			queue = append(queue, Node{node.Child, node.Child.Left}, Node{node.Child, node.Child.Right})
		}

		if xParent != nil || yParent != nil {
			return xParent != yParent && xParent != nil && yParent != nil
		}
	}
	return false
}
