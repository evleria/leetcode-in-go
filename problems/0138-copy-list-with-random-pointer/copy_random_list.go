package _138_copy_list_with_random_pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(node *Node) *Node {
	return deepClone(node, map[*Node]*Node{})
}

func deepClone(node *Node, m map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	if n, ok := m[node]; ok {
		return n
	}

	newNode := &Node{
		Val: node.Val,
	}
	m[node] = newNode
	newNode.Next, newNode.Random = deepClone(node.Next, m), deepClone(node.Random, m)
	return newNode
}
