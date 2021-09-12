package _882_reachable_nodes_in_subdivided_graph

import (
	"container/heap"
)

type MaxHeap [][2]int

func (h *MaxHeap) Less(i, j int) bool { return (*h)[i][1] > (*h)[j][1] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
func (h *MaxHeap) Push(v interface{}) { *h = append(*h, v.([2]int)) }

func reachableNodes(edges [][]int, maxMoves int, n int) int {
	matrix := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make(map[int]int)
	}

	for _, edge := range edges {
		matrix[edge[0]][edge[1]] = edge[2]
		matrix[edge[1]][edge[0]] = edge[2]
	}

	result, visited := 0, make([]bool, n)
	h := new(MaxHeap)
	heap.Push(h, [2]int{0, maxMoves})

	for h.Len() != 0 {
		cur := heap.Pop(h).([2]int)
		curNode, moves := cur[0], cur[1]
		if visited[curNode] {
			continue
		}
		visited[curNode] = true
		result++
		for i, v := range matrix[curNode] {
			cost := v + 1
			if cost > 0 {
				if moves >= cost && !visited[i] {
					heap.Push(h, [2]int{i, moves - cost})
				}
				reach := min(moves, cost-1)
				result += reach
				matrix[i][curNode] -= reach
			}
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
