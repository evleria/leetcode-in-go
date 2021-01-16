package _152_maximum_product_subarray

import "math"

func maxProduct(nums []int) int {
	min, max, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			min, max = max, min
		}
		min = int(math.Min(float64(nums[i]), float64(min*nums[i])))
		max = int(math.Max(float64(nums[i]), float64(max*nums[i])))
		if max > result {
			result = max
		}
	}
	return result
}
