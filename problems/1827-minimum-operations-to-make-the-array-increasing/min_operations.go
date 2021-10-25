package _827_minimum_operations_to_make_the_array_increasing

func minOperations(nums []int) int {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			count += nums[i-1] - nums[i] + 1
			nums[i] = nums[i-1] + 1
		}
	}
	return count
}
