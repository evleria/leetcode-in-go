package _470_shuffle_the_array

func shuffle(nums []int, n int) []int {
	first, second := nums[:n], nums[n:]
	result := make([]int, 0, len(nums))
	for i := 0; i < len(first); i++ {
		result = append(result, first[i], second[i])
	}
	return result
}
