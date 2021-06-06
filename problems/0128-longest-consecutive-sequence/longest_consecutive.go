package _128_longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	m := map[int]int{}
	res := 0
	for _, n := range nums {
		if _, ok := m[n]; !ok {
			left, right := m[n-1], m[n+1]
			l := left + right + 1
			m[n] = l
			if left > 0 {
				m[n-left] = l
			}
			if right > 0 {
				m[n+right] = l
			}
			if l > res {
				res = l
			}
		}
	}

	return res
}
