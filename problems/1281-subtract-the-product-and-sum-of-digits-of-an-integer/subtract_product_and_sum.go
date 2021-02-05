package _281_subtract_the_product_and_sum_of_digits_of_an_integer

func subtractProductAndSum(n int) int {
	product, sum := 1, 0
	for ; n != 0; n /= 10 {
		d := n % 10
		product, sum = product*d, sum+d
	}
	return product - sum
}
