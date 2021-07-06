package _338_reduce_array_size_to_the_half

import "sort"

func minSetSize(arr []int) int {
	dict := map[int]int{}
	for _, n := range arr {
		dict[n]++
	}

	counters := make([]int, 0, len(dict))
	for _, v := range dict {
		counters = append(counters, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counters)))

	half := len(arr) / 2
	for i, n := range counters {
		half -= n
		if half <= 0 {
			return i + 1
		}
	}

	return 0
}
