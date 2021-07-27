package _913_maximum_product_difference_between_two_pairs

import "sort"

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]*nums[len(nums)-2] - nums[0]*nums[1]
}
