package _989_add_to_array_form_of_integer

func addToArrayForm(num []int, k int) []int {
	for i := len(num) - 1; i >= 0 && k > 0; i-- {
		num[i] += k
		k, num[i] = num[i]/10, num[i]%10
	}

	for ; k > 0; k /= 10 {
		num = append([]int{k % 10}, num...)
	}

	return num
}
