package _970_powerful_integers

func powerfulIntegers(x int, y int, bound int) []int {
	m := map[int]bool{}
	for a := 1; a < bound; a *= x {
		for b := 1; a+b <= bound; b *= y {
			m[a+b] = true
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}

	result := []int{}
	for n := range m {
		result = append(result, n)
	}
	return result
}
