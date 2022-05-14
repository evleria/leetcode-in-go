package _743_network_delay_time

func max(dis []int) (md int) {
	md = dis[0]
	for _, d := range dis {
		if d < 0 {
			return -1
		}
		if d > md {
			md = d
		}
	}
	return
}
func min(dis []int, solved []bool) (ind int) {
	ind = -1
	mv := -1
	for i, d := range dis {
		if !solved[i] && l(d, mv) {
			mv = d
			ind = i
		}
	}
	return
}
func l(v1, v2 int) bool {
	if v1 < 0 {
		return false
	}
	if v2 < 0 {
		return true
	}
	return v1 < v2
}
func add(v1, v2 int) int {
	if v1 == -1 || v2 == -1 {
		return -1
	}
	return v1 + v2
}
func networkDelayTime(times [][]int, n, k int) int {
	var solved = make([]bool, n)
	k -= 1
	solved[k] = true
	var graph = make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, n)
		for j := range graph[i] {
			graph[i][j] = -1
		}
		graph[i][i] = 0
	}
	for _, time := range times {
		graph[time[0]-1][time[1]-1] = time[2]
	}
	for _ = range make([]byte, n-1) {
		cind := min(graph[k], solved)
		if cind == -1 {
			return -1
		}
		solved[cind] = true
		for i := range make([]byte, n) {
			if i == k || i == cind {
				continue
			}
			newDis := add(graph[k][cind], graph[cind][i])
			if l(newDis, graph[k][i]) {
				graph[k][i] = newDis
			}
		}
	}
	return max(graph[k])
}
