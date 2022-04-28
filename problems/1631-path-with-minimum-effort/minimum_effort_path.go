package _631_path_with_minimum_effort

import (
	"container/heap"
	"math"
)

type minHeap []Cost

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i].effort < m[j].effort }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(Cost))
}
func (m *minHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

type Cell struct {
	r, c int
}

type Cost struct {
	Cell
	effort int
}

var dirs = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func minimumEffortPath(heights [][]int) int {
	e := make([][]int, len(heights))
	for i := range heights {
		e[i] = make([]int, len(heights[0]))
	}
	// to get anywhere from 0,0 it is infinite cost
	for i := range heights {
		for j := range heights[0] {
			e[i][j] = math.MaxInt64
		}
	}
	e[0][0] = 0

	q := minHeap{}
	heap.Push(&q, Cost{Cell{0, 0}, 0})
	rl, cl := len(heights), len(heights[0])

	for len(q) > 0 {
		cur := heap.Pop(&q).(Cost)

		if cur.r == rl-1 && cur.c == cl-1 {
			return cur.effort
		}

		for _, d := range dirs {
			nextR, nextC := cur.r+d[0], cur.c+d[1]
			if nextR < 0 || nextC < 0 || nextR > rl-1 || nextC > cl-1 {
				continue
			}
			// the cost to the next cell is either max effort to get here currently
			// or getting from cur cell to next cell
			nextCell := Cost{Cell{nextR, nextC}, max(cur.effort, abs(heights[cur.r][cur.c]-heights[nextR][nextC]))}
			// if existing effort to get to next cell > newly calculated effort, add to queue of cells to try next
			if e[nextR][nextC] > nextCell.effort {
				heap.Push(&q, nextCell)
				e[nextR][nextC] = nextCell.effort
			}
		}
	}
	return -1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
