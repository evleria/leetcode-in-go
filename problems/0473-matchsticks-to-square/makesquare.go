package _473_matchsticks_to_square

import "sort"

func makesquare(matchsticks []int) bool {
	per := 0
	for _, value := range matchsticks {
		per += value
	}
	if per%4 != 0 {
		return false
	}

	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))

	edge := [4]int{}
	return dfs(matchsticks, 0, edge, per/4)
}

func dfs(a []int, index int, edge [4]int, maxEdge int) bool {
	if index == len(a) {
		return edge[0] == edge[1] && edge[1] == edge[2] && edge[2] == edge[3]
	}

	for i := 0; i < 4; i++ {
		if edge[i]+a[index] > maxEdge {
			continue
		}
		edge[i] += a[index]
		if dfs(a, index+1, edge, maxEdge) {
			return true
		}
		edge[i] -= a[index]
	}
	return false
}
