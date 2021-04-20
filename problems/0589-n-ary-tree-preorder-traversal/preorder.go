package _589_n_ary_tree_preorder_traversal

type Node struct {
	Val      int
	Children []*Node
}

//goland:noinspection GoUnusedFunction
func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	var dfs func(*Node)
	dfs = func(node *Node) {
		result = append(result, node.Val)
		for _, child := range node.Children {
			dfs(child)
		}
	}

	dfs(root)

	return result
}
