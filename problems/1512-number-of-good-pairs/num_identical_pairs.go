package _512_number_of_good_pairs

func numIdenticalPairs(nums []int) int {
	m := [100]int{}
	result := 0
	for i := 0; i < len(nums); i++ {
		result += m[nums[i]-1]
		m[nums[i]-1]++
	}
	return result
}
