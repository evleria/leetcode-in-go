package _133_clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
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
	for _, neighbor := range node.Neighbors {
		neighbor.Neighbors = append(neighbor.Neighbors, deepClone(neighbor, m))
	}
	return newNode
}
