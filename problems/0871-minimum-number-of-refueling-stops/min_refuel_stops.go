package _871_minimum_number_of_refueling_stops

import "container/heap"

type MaxHeap []int

func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *MaxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MaxHeap) Len() int           { return len(*h) }
func (h *MaxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}
func (h *MaxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }

func minRefuelStops(target int, fuel int, stations [][]int) int {
	h, pos, count := new(MaxHeap), 0, 0

	tryFuelUpRetrospectively := func(distance int) bool {
		fuel -= distance
		for fuel < 0 && h.Len() > 0 {
			fuel += heap.Pop(h).(int)
			count++
		}
		return fuel >= 0
	}

	for _, station := range stations {
		if !tryFuelUpRetrospectively(station[0] - pos) {
			return -1
		}

		heap.Push(h, station[1])
		pos = station[0]
	}

	if !tryFuelUpRetrospectively(target - pos) {
		return -1
	}

	return count
}
