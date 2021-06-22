package _389_create_target_array_in_the_given_order

func createTargetArray(nums []int, index []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(index); i++ {
		n, idx := nums[i], index[i]
		for j := i - 1; j >= idx; j-- {
			result[j+1] = result[j]
		}
		result[idx] = n
	}
	return result
}
