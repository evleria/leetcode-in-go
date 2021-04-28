package _977_squares_of_a_sorted_array

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	left, right := 0, len(nums)-1
	for i := len(result) - 1; i >= 0; i-- {
		if -1*nums[left] < nums[right] {
			result[i] = nums[right] * nums[right]
			right--
		} else {
			result[i] = nums[left] * nums[left]
			left++
		}
	}
	return result
}
