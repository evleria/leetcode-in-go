package _257_binary_tree_paths

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	var result []string
	dfs(root, nil, &result)
	return result
}

func dfs(node *TreeNode, path []string, result *[]string) {
	if node == nil {
		return
	}
	path = append(path, strconv.Itoa(node.Val))
	if node.Left == nil && node.Right == nil {
		*result = append(*result, strings.Join(path, "->"))
	} else {
		dfs(node.Left, path, result)
		dfs(node.Right, path, result)
	}
}
