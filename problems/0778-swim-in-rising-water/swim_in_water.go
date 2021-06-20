package _778_swim_in_rising_water

import "container/heap"

type Node struct {
	Elevation int
	X, Y      int
}

type MinHeap []Node

func (h *MinHeap) Less(i, j int) bool { return (*h)[i].Elevation < (*h)[j].Elevation }
func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
func (h *MinHeap) Push(v interface{}) { *h = append(*h, v.(Node)) }

var directions = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func swimInWater(grid [][]int) int {
	n := len(grid)
	visited := make([]bool, n*n)
	var t int
	h := new(MinHeap)
	heap.Push(h, Node{grid[0][0], 0, 0})

	var dfs func(x, y int) bool
	dfs = func(x, y int) bool {
		if x == n-1 && y == n-1 {
			return true
		}

		visited[grid[x][y]] = true

		for _, dir := range directions {
			x0, y0 := x+dir[0], y+dir[1]
			if x0 >= 0 && x0 < n && y0 >= 0 && y0 < n && !visited[grid[x0][y0]] {
				if e := grid[x0][y0]; e <= t {
					if r := dfs(x0, y0); r {
						return r
					}
				} else {
					heap.Push(h, Node{e, x0, y0})
				}
			}
		}

		return false
	}

	for {
		node := heap.Pop(h).(Node)
		t = node.Elevation

		if r := dfs(node.X, node.Y); r {
			return t
		}
	}
}
