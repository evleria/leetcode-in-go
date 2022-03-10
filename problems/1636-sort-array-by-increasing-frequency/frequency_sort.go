package _636_sort_array_by_increasing_frequency

import "sort"

type entry struct{ num, count int }

func frequencySort(nums []int) []int {
	m := map[int]int{}
	for _, n := range nums {
		m[n]++
	}
	entries := make([]entry, 0, len(m))
	for k, v := range m {
		entries = append(entries, entry{num: k, count: v})
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].count != entries[j].count {
			return entries[i].count < entries[j].count
		}
		return entries[i].num > entries[j].num
	})
	result := make([]int, 0, len(nums))
	for _, e := range entries {
		for i := 0; i < e.count; i++ {
			result = append(result, e.num)
		}
	}
	return result
}
