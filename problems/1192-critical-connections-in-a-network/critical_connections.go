package _192_critical_connections_in_a_network

func criticalConnections(n int, connections [][]int) [][]int {
	graph := make([][]int, n)
	for _, c := range connections {
		graph[c[0]] = append(graph[c[0]], c[1])
		graph[c[1]] = append(graph[c[1]], c[0])
	}
	result := [][]int{}
	dfs(0, -1, 0, graph, make([]*int, n), &result)
	return result
}

func dfs(cur int, prev int, depth int, graph [][]int, visited []*int, result *[][]int) {
	visited[cur] = &depth
	for _, n := range graph[cur] {
		if n == prev {
			continue
		}
		if visited[n] == nil {
			dfs(n, cur, depth+1, graph, visited, result)
		}
		if *visited[cur] > *visited[n] {
			visited[cur] = visited[n]
		}
		if *visited[n] > depth {
			*result = append(*result, []int{cur, n})
		}
	}
}
