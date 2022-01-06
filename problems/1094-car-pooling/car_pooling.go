package _094_car_pooling

func carPooling(trips [][]int, capacity int) bool {
	diff := [1001]int{}
	start, end := 1000, 0
	for _, trip := range trips {
		p, from, to := trip[0], trip[1], trip[2]
		diff[from] += p
		diff[to] -= p

		if start > from {
			start = from
		}
		if end < to {
			end = to
		}
	}

	for p, i := 0, start; i < end; i++ {
		p += diff[i]
		if p > capacity {
			return false
		}
	}

	return true
}
