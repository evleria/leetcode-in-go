package _160_minimum_sum_of_four_digit_number_after_splitting_digits

import "sort"

func minimumSum(num int) int {
	d := []int{}
	for num > 0 {
		d = append(d, num%10)
		num /= 10
	}
	sort.Ints(d)
	new1 := d[0]*10 + d[2]
	new2 := d[1]*10 + d[3]
	return new1 + new2
}
