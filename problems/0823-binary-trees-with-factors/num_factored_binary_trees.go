package _823_binary_trees_with_factors

import "sort"

func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	m := make(map[int]int, len(arr))
	result := 0
	for _, n := range arr {
		count := 1
		for k, v := range m {
			if n%k == 0 {
				count += v * m[n/k]
			}
		}
		m[n] = count
		result += count
	}
	return result % 1000000007
}
