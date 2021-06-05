package _383_maximum_performance_of_a_team

import (
	"container/heap"
	"sort"
)

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

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	candidates := make([][2]int, n)
	for i := 0; i < n; i++ {
		candidates[i] = [2]int{efficiency[i], speed[i]}
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i][0] > candidates[j][0]
	})

	maxPerf, speedSum, h := 0, 0, new(MinHeap)
	for _, candidate := range candidates {
		if h.Len() >= k {
			speedSum -= heap.Pop(h).(int)
		}
		speedSum += candidate[1]
		heap.Push(h, candidate[1])

		if perf := speedSum * candidate[0]; perf > maxPerf {
			maxPerf = perf
		}
	}

	return maxPerf % 1_000_000_007
}
