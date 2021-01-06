package _016_3Sum_closest

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	result := nums[0] + nums[1] + nums[2]
	for start := 0; start < len(nums)-2; start++ {
		left, right := start+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right] + nums[start]
			if abs(result-target) > abs(sum-target) {
				result = sum
			}

			switch {
			case sum < target:
				left++
			case sum > target:
				right--
			default:
				return target
			}
		}
	}

	return result
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
