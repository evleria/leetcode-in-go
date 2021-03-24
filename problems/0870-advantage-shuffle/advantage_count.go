package _870_advantage_shuffle

import "sort"

func advantageCount(A []int, B []int) []int {
	sort.Ints(A)
	b := make([][2]int, len(B))
	for i, n := range B {
		b[i] = [2]int{n, i}
	}
	sort.Slice(b, func(i, j int) bool {
		return b[i][0] < b[j][0]
	})

	result := make([]int, len(A))
	left, right := 0, len(A)-1
	for i := len(b) - 1; i >= 0; i-- {
		val, idx := b[i][0], b[i][1]
		if A[right] > val {
			result[idx] = A[right]
			right--
		} else {
			result[idx] = A[left]
			left++
		}
	}
	return result
}
