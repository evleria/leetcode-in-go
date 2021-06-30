package _566_detect_pattern_of_length_m_repeated_k_or_more_times

func containsPattern(arr []int, m int, k int) bool {
	isPattern := func(i int) bool {
		for j := i; j < m+i; j++ {
			for z := 1; z < k; z++ {
				if arr[j] != arr[j+z*m] {
					return false
				}
			}
		}
		return true
	}

	for i := 0; i < len(arr)-m*k+1; i++ {
		if isPattern(i) {
			return true
		}
	}
	return false
}
