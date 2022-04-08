package _703_kth_largest_element_in_a_stream

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

type KthLargest struct {
	k    int
	nums *MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	result := KthLargest{k: k, nums: &MinHeap{}}
	for _, num := range nums {
		result.add(num)
	}

	return result
}

func (k *KthLargest) Add(val int) int {
	k.add(val)
	return k.nums.Peak()
}

func (k *KthLargest) add(val int) {
	heap.Push(k.nums, val)
	if k.nums.Len() > k.k {
		heap.Pop(k.nums)
	}
}
