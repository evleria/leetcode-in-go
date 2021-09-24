package _137_n_th_tribonacci_number

func tribonacci(n int) int {
	if n < 2 {
		return n
	}

	a, b, c := 0, 1, 1

	for i := 3; i <= n; i++ {
		d := a + b + c
		a, b, c = b, c, d
	}

	return c
}
