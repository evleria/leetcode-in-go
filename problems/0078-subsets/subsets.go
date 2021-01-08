package _078_subsets

func subsets(nums []int) [][]int {
	result := make([][]int, 0, 16)
	buffer := make([]int, 0, 16)
	backtrack(&result, buffer, nums)
	return result
}

func backtrack(result *[][]int, buffer []int, nums []int) {
	*result = append(*result, append([]int{}, buffer...))
	for i, n := range nums {
		backtrack(result, append(buffer, n), nums[i+1:])
	}
}
