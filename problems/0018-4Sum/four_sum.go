package _018_4Sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}

	sort.Ints(nums)
	result := [][]int{}
	for first := 0; first < len(nums)-3; first++ {
		for second := first + 1; second < len(nums)-2; second++ {
			left, right := second+1, len(nums)-1
			for left < right {
				sum := nums[first] + nums[second] + nums[left] + nums[right]
				switch {
				case sum < target:
					left++
				case sum > target:
					right--
				default:
					result = append(result, []int{nums[first], nums[second], nums[left], nums[right]})
					for ; left < right && nums[left] == nums[left+1]; left++ {
					}
					for ; right > left && nums[right] == nums[right-1]; right-- {
					}
					left++
					right--
				}
			}
			for ; second < len(nums)-3 && nums[second] == nums[second+1]; second++ {
			}
		}
		for ; first < len(nums)-4 && nums[first] == nums[first+1]; first++ {
		}
	}

	return result
}
