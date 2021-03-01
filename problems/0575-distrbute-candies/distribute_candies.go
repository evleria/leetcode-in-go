package _575_distrbute_candies

func distributeCandies(candyType []int) int {
	max := len(candyType) / 2
	m := map[int]bool{}
	for _, t := range candyType {
		m[t] = true
		if len(m) == max {
			break
		}
	}

	return len(m)
}
