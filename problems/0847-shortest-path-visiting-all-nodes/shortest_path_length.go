package _847_shortest_path_visiting_all_nodes

func shortestPathLength(graph [][]int) int {
	if len(graph) == 1 {
		return 0
	}

	visited := make([]map[int]bool, len(graph))
	queue := make([][]int, 0)
	for i := range graph {
		mask := 1 << i
		visited[i] = map[int]bool{}
		queue = append(queue, []int{i, mask})
	}

	target := 1<<len(graph) - 1
	for pathLen := 1; ; pathLen++ {
		for _, cur := range queue {
			queue = queue[1:]
			for _, neigh := range graph[cur[0]] {
				nextMask := cur[1] | (1 << neigh)
				if !visited[neigh][nextMask] {
					if nextMask == target {
						return pathLen
					}
					visited[neigh][nextMask] = true
					queue = append(queue, []int{neigh, nextMask})
				}
			}
		}
	}
}
