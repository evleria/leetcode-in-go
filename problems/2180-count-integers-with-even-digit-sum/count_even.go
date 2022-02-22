package _180_count_integers_with_even_digit_sum

func countEven(num int) int {
	result := 0
	for i := 1; i <= num; i++ {
		n, sum := i, 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		if sum%2 == 0 {
			result++
		}
	}

	return result
}
