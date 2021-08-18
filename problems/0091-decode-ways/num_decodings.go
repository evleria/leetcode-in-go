package _091_decode_ways

func numDecodings(s string) int {
	a, b := 1, 1
	if s[len(s)-1] == '0' {
		a = 0
	}
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == '0' {
			a, b = 0, a
		} else if s[i] == '1' || (s[i] == '2' && s[i+1] <= '6') {
			a, b = a+b, a
		} else {
			b = a
		}
	}
	return a
}
