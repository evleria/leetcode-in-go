package _116_populating_next_right_pointers_in_each_node

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

	queue := []*Node{root}
	for len(queue) > 0 {
		var next *Node
		for _, node := range queue {
			queue = queue[1:]
			node.Next, next = next, node
			if node.Left != nil {
				queue = append(queue, node.Right, node.Left)
			}
		}
	}
	return root
}
