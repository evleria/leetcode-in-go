package _402_remove_k_digits

func removeKdigits(num string, k int) string {
	r, n := []byte(num), 0
	for _, c := range r {
		for n > 0 && r[n-1] > c && k > 0 {
			n, k = n-1, k-1
		}
		if c != '0' || n > 0 {
			r[n] = c
			n++
		}
	}
	n -= k

	if n <= 0 {
		return "0"
	}

	return string(r[:n])
}
