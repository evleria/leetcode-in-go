package _675_minimize_deviation_in_array

import (
	"container/heap"
	"math"
)

type MaxHeap []int

func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
func (h *MaxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }

func minimumDeviation(nums []int) int {
	h := new(MaxHeap)
	m := math.MaxInt64

	for _, num := range nums {
		if num%2 != 0 {
			num *= 2
		}
		heap.Push(h, num)
		m = min(m, num)
	}

	deviation := math.MaxInt64
	for h.Len() > 0 {
		num := heap.Pop(h).(int)
		deviation = min(deviation, num-m)

		if num%2 != 0 {
			break
		}
		heap.Push(h, num/2)
		m = min(m, num/2)
	}

	return deviation
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
