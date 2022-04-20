package _173_binary_search_tree_iterator

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	var stack []*TreeNode
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}
	return BSTIterator{stack}
}

func (b *BSTIterator) Next() int {
	node := b.stack[len(b.stack)-1]
	val := node.Val
	b.stack = b.stack[:len(b.stack)-1]
	if node.Right != nil {
		node = node.Right
		for node != nil {
			b.stack = append(b.stack, node)
			node = node.Left
		}
	}

	return val
}

func (b *BSTIterator) HasNext() bool {
	return len(b.stack) > 0
}
