package _430_flatten_a_multilevel_doubly_linked_list

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	helper(root)
	return root
}

func helper(node *Node) *Node {
	var prev *Node
	for ; node != nil; node = node.Next {
		if node.Child != nil {
			last := helper(node.Child)
			if node.Next != nil {
				node.Next.Prev = last
			}
			node.Child, node.Next, node.Child.Prev, last.Next = nil, node.Child, node, node.Next
		}

		prev = node
	}
	return prev
}
