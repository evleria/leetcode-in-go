package _639_decode_ways_II

func numDecodings(s string) int {
	if len(s) > 0 && s[0] == '0' {
		return 0
	}

	for i, r := range s {
		if r == '*' {
			sum := 0
			for ch := '1'; ch <= '9'; ch++ {
				sum += helper(s, i, ch)
			}
			return sum
		}
		return helper(s, i, r)
	}

	return 1
}

func helper(s string, i int, r rune) int {
	switch r {
	case '1':
		if i < len(s)-1 {
			return numDecodings(s[i+1:]) + numDecodings(s[i+2:])
		}

	case '2':
		if i < len(s)-1 && s[i+1] <= '6' {
			return numDecodings(s[i+1:]) + numDecodings(s[i+2:])
		}
	}

	return numDecodings(s[i+1:])
}
