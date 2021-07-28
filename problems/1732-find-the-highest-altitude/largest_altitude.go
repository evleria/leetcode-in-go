package _732_find_the_highest_altitude

func largestAltitude(gain []int) int {
	n, max := 0, 0
	for i := 0; i < len(gain); i++ {
		n += gain[i]
		if n > max {
			max = n
		}
	}
	return max
}
