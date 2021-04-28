package _837_sum_of_digits_in_base_k

func sumBase(n int, k int) int {
	sum := 0
	for ; n > 0; n /= k {
		sum += n % k
	}
	return sum
}
