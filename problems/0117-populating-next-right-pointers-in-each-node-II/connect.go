package _117_populating_next_right_pointers_in_each_node_II

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var queue []*Node
	queue = append(queue, root)
	for len(queue) > 0 {
		var prev *Node
		for _, node := range queue {
			queue = queue[1:]
			node.Next, prev = prev, node
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
	}
	return root
}
