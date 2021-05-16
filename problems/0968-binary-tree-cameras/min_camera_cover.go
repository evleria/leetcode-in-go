package _968_binary_tree_cameras

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

const (
	NeedToCover = iota
	NoNeedToCover
	Camera
)

func minCameraCover(root *TreeNode) int {
	cameras := 0
	if dfs(root, &cameras) == NeedToCover {
		cameras++
	}
	return cameras
}

func dfs(node *TreeNode, cameras *int) int {
	if node == nil {
		return NoNeedToCover
	}

	left, right := dfs(node.Left, cameras), dfs(node.Right, cameras)

	switch {
	case left == NeedToCover || right == NeedToCover:
		*cameras++
		return Camera
	case left == NoNeedToCover && right == NoNeedToCover:
		return NeedToCover
	default:
		return NoNeedToCover
	}
}
