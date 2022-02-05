package _023_merge_k_sorted_lists

import (
	"container/heap"
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

type minHeap []*ListNode

func (h *minHeap) Less(i, j int) bool { return (*h)[i].Val < (*h)[j].Val }
func (h *minHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *minHeap) Len() int           { return len(*h) }
func (h *minHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
func (h *minHeap) Push(v interface{}) { *h = append(*h, v.(*ListNode)) }

func mergeKLists(lists []*ListNode) *ListNode {
	h := new(minHeap)
	for _, head := range lists {
		if head != nil {
			heap.Push(h, head)
		}
	}

	if len(*h) == 0 {
		return nil
	}

	dummy := new(ListNode)
	cur := dummy
	for h.Len() > 1 {
		node := heap.Pop(h).(*ListNode)
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
		cur, cur.Next = node, node
	}
	cur.Next = (*h)[0]

	return dummy.Next
}
