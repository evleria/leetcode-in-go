package _929_concatenation_of_array

func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, n*2)
	for i := 0; i < n; i++ {
		ans[i], ans[i+n] = nums[i], nums[i]
	}
	return ans
}
