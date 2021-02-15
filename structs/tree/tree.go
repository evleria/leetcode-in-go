package tree

import "math"

const NULL = math.MinInt64

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func FromSlice(input []int) *TreeNode {
	if len(input) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: input[0],
	}

	queue := make([]*TreeNode, 1, len(input))
	queue[0] = root

	for i := 1; i < len(input); {
		node := queue[0]
		queue = queue[1:]

		if i < len(input) && input[i] != NULL {
			node.Left = &TreeNode{Val: input[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(input) && input[i] != NULL {
			node.Right = &TreeNode{Val: input[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func (root *TreeNode) ToSlice() []int {
	if root == nil {
		return []int{}
	}

	queue, result := make([]*TreeNode, 1), make([]int, 0, 16)

	queue[0] = root
	l := 1

	for len(queue) != 0 {
		if cur := queue[0]; cur != nil {
			result = append(result, cur.Val)
			l = len(result)
			queue = append(queue, cur.Left, cur.Right)
		} else {
			result = append(result, NULL)
		}
		queue = queue[1:]
	}

	return result[:l]
}

func (root *TreeNode) FindInBST(n int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val == n:
		return root
	case root.Val > n:
		return root.Left.FindInBST(n)
	case root.Val < n:
		return root.Right.FindInBST(n)
	}

	return nil
}

func (root *TreeNode) FindInBT(n int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == n {
		return root
	}
	if left := root.Left.FindInBT(n); left != nil {
		return left
	}
	return root.Right.FindInBT(n)
}
