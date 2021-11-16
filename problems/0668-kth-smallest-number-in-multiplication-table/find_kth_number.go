package _668_kth_smallest_number_in_multiplication_table

import "sort"

func findKthNumber(m int, n int, k int) int {
	if m > n {
		return findKthNumber(n, m, k)
	}

	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = i + 1
	}

	res := sort.Search(n*m+1, func(i int) bool {
		index := 0
		a := 1
		for a <= m {
			temp := sort.Search(n, func(j int) bool {
				return list[j]*a > i
			})
			if temp == 0 {
				break
			} else {
				a++
			}
			index += temp
		}
		return index >= k
	})

	return res
}
