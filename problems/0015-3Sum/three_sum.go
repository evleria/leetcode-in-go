package _015_3Sum

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	result := make([][]int, 0, 4)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			switch {
			case sum == 0:
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for ; left < len(nums)-1 && nums[left] == nums[left+1]; left++ {
				}
				for ; right > 0 && nums[right] == nums[right-1]; right-- {
				}

				left++
				right--
			case sum > 0:
				right--
			case sum < 0:
				left++
			}
		}

		for ; i < len(nums)-1 && nums[i] == nums[i+1]; i++ {
		}
	}

	return result
}
