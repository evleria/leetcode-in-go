package _785_is_graph_bipartite

func isBipartite(graph [][]int) bool {
	n := len(graph)
	color := make([]int, n)
	queue := make([]int, 0)

	for i := 0; i < n; i++ {
		if color[i] == 0 {
			color[i] = 1
			queue = append(queue, i)
			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]
				for _, v := range graph[cur] {
					if color[v] == color[cur] {
						return false
					}
					if color[v] == 0 {
						color[v] = -color[cur]
						queue = append(queue, v)
					}
				}
			}
		}
	}
	return true
}
