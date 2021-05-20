package _295_find_numbers_with_even_number_of_digits

func findNumbers(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		digits := 0
		for nums[i] != 0 {
			nums[i] /= 10
			digits++
		}
		if digits%2 == 0 {
			count++
		}

	}

	return count
}
