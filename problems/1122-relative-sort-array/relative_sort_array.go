package _122_relative_sort_array

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	order := make(map[int]int, len(arr2))
	for i, n := range arr2 {
		order[n] = i
	}

	sort.Slice(arr1, func(i, j int) bool {
		ai, oki := order[arr1[i]]
		aj, okj := order[arr1[j]]

		switch {
		case oki && okj:
			return ai < aj
		case !oki && !okj:
			return arr1[i] < arr1[j]
		default:
			return oki
		}
	})

	return arr1
}
