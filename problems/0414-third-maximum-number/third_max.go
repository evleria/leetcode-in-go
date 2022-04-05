package _414_third_maximum_number

import "math"

func thirdMax(nums []int) int {
	if len(nums) < 3 {
		return maximum(nums)
	}
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] == max1 || nums[i] == max2 || nums[i] == max3 {
			continue
		}
		if nums[i] > max1 {
			max3 = max2
			max2 = max1
			max1 = nums[i]
		} else if nums[i] > max2 {
			max3 = max2
			max2 = nums[i]
		} else if nums[i] > max3 {
			max3 = nums[i]
		}
	}
	if max3 == math.MinInt64 {
		return max1
	}
	return max3
}

func maximum(nums []int) int {
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
