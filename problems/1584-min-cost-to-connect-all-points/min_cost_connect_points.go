package _584_min_cost_to_connect_all_points

import "sort"

type point struct {
	x    int
	y    int
	dist int
}

func minCostConnectPoints(points [][]int) int {
	distances := []point{}
	for p1 := 0; p1 < len(points); p1++ {
		for p2 := p1 + 1; p2 < len(points); p2++ {
			x, y := points[p1], points[p2]
			distances = append(distances, point{p1, p2, getDistance(x, y)})
		}
	}

	parent := make([]int, len(points))
	for i := range parent {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(x, y int) bool {
		px := find(x)
		py := find(y)
		if px != py {
			parent[px] = py
			return true
		}
		return false
	}

	sort.Slice(distances, func(x, y int) bool {
		return distances[x].dist < distances[y].dist
	})

	cost := 0
	for _, edge := range distances {
		if union(edge.x, edge.y) {
			cost += edge.dist
		}
	}

	return cost
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getDistance(x, y []int) int {
	return abs(x[0]-y[0]) + abs(x[1]-y[1])
}
