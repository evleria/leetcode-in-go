package _202_smallest_string_with_swaps

import "sort"

func smallestStringWithSwaps(s string, pairs [][]int) string {
	graph := make([][]int, len(s))
	for _, p := range pairs {
		graph[p[0]] = append(graph[p[0]], p[1])
		graph[p[1]] = append(graph[p[1]], p[0])
	}
	res := []byte(s)
	visited := make([]bool, len(s))
	for i := range s {
		if !visited[i] {
			cc := bfs(graph, i, visited)
			var list []byte
			for _, v := range cc {
				list = append(list, s[v])
			}
			sort.Slice(list, func(i, j int) bool {
				return list[i] < list[j]
			})
			for i, v := range cc {
				res[v] = list[i]
			}
		}
	}
	return string(res)
}

func bfs(graph [][]int, n int, visited []bool) []int {
	res := []int{n}
	visited[n] = true
	queue := []int{n}
	for len(queue) > 0 {
		for i := 0; i < len(queue); i++ {
			for _, v := range graph[queue[i]] {
				if !visited[v] {
					visited[v] = true
					res = append(res, v)
					queue = append(queue, v)
				}
			}
		}
		queue = queue[len(queue):]
	}
	sort.Ints(res)
	return res
}
