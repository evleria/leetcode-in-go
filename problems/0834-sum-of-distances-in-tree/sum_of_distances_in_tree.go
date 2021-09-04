package _834_sum_of_distances_in_tree

func sumOfDistancesInTree(n int, edges [][]int) []int {
	adjacent := make([][]int, n)
	for _, edge := range edges {
		adjacent[edge[0]] = append(adjacent[edge[0]], edge[1])
		adjacent[edge[1]] = append(adjacent[edge[1]], edge[0])
	}

	res, count := make([]int, n), make([]int, n)

	var postDfs func(node, parent int)
	postDfs = func(node, parent int) {
		for _, dest := range adjacent[node] {
			if dest != parent {
				postDfs(dest, node)
				res[node] += res[dest] + count[dest]
				count[node] += count[dest]
			}
		}
		count[node]++
	}

	var preDfs func(node, parent int)
	preDfs = func(node, parent int) {
		for _, dest := range adjacent[node] {
			if dest != parent {
				res[dest] = res[node] - count[dest] + n - count[dest]
				preDfs(dest, node)
			}
		}
	}

	postDfs(0, -1)
	preDfs(0, -1)

	return res
}
