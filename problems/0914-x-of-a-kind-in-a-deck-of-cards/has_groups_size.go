package _914_x_of_a_kind_in_a_deck_of_cards

func hasGroupsSizeX(deck []int) bool {
	m := map[int]int{}
	for _, n := range deck {
		m[n]++
	}

	x := m[deck[0]]
	for _, c := range m {
		x = gcd(x, c)
		if x == 1 {
			return false
		}
	}

	return true
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
