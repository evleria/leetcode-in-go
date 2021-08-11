package _954_array_of_doubled_pairs

import "sort"

func canReorderDoubled(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool {
		return abs(arr[i]) < abs(arr[j])
	})

	m := map[int]int{}
	for _, n := range arr {
		if x, ok := m[n/2]; ok && x != 0 && n%2 == 0 {
			m[n/2]--
		} else {
			m[n]++
		}
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
