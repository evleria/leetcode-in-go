package _429_n_ary_tree_level_order_traversal

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*Node{root}

	for len(queue) > 0 {
		var level []int
		for _, node := range queue {
			queue = queue[1:]
			level = append(level, node.Val)
			queue = append(queue, node.Children...)
		}
		result = append(result, level)
	}

	return result
}
