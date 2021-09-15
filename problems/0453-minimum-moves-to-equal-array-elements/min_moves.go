package _453_minimum_moves_to_equal_array_elements

func minMoves(nums []int) int {
	min, sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] < min {
			min = nums[i]
		}
	}
	return sum - len(nums)*min
}
