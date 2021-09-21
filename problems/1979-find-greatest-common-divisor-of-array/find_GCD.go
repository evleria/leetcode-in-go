package _979_find_greatest_common_divisor_of_array

func findGCD(nums []int) int {
	min, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		} else if nums[i] > max {
			max = nums[i]
		}
	}

	return gcd(max, min)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
