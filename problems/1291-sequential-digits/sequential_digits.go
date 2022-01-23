package _291_sequential_digits

func sequentialDigits(low int, high int) []int {
	result := []int{}
	seed, increment := 12, 11
	for i := 2; i < 10; i++ {
		row := generate(seed, increment, i)
		for _, n := range row {
			if n > high {
				return result
			} else if n >= low {
				result = append(result, n)
			}
		}
		seed, increment = seed*10+i+1, increment*10+1
	}

	return result
}

func generate(seed, increment, digits int) []int {
	result := make([]int, 0, 10-digits)
	result = append(result, seed)
	for i := 0; i < 10-digits-1; i++ {
		seed += increment
		result = append(result, seed)
	}
	return result
}
