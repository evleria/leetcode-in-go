package _684_redundant_connection

func findRedundantConnection(edges [][]int) []int {
	m := len(edges)
	adjacencyList := make([][][2]int, m)
	for i, edge := range edges {
		from, to := edge[0], edge[1]
		adjacencyList[from-1] = append(adjacencyList[from-1], [2]int{to - 1, i})
		adjacencyList[to-1] = append(adjacencyList[to-1], [2]int{from - 1, i})
	}

	visited := make(map[int]int, m)
	var dfs func(to, from, step int, history []int) *int
	dfs = func(to, from, step int, history []int) *int {
		visited[to] = step
		for _, node := range adjacencyList[to] {
			dest := node[0]
			if dest == from {
				continue
			}

			if prevStep, ok := visited[dest]; ok {
				history = append(history, node[1])
				max := 0
				for _, i := range history[prevStep:] {
					if i > max {
						max = i
					}
				}
				return &max
			}

			if r := dfs(dest, to, step+1, append(history, node[1])); r != nil {
				return r
			}
		}
		delete(visited, to)

		return nil
	}

	i := dfs(0, -1, 0, nil)
	return edges[*i]
}
