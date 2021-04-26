package _642_furthest_building_you_can_reach

import "container/heap"

type MinHeap []int

func (h *MinHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MinHeap) Len() int           { return len(*h) }
func (h *MinHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
func (h *MinHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *MinHeap) Peak() int          { return (*h)[0] }

func furthestBuilding(heights []int, bricks int, ladders int) int {
	var h = &MinHeap{}
	for i := 0; i < len(heights)-1; i++ {
		diff := heights[i+1] - heights[i]
		if diff > 0 {
			heap.Push(h, diff)
		}
		if h.Len() > ladders {
			bricks -= heap.Pop(h).(int)
		}
		if bricks < 0 {
			return i
		}
	}

	return len(heights) - 1
}
