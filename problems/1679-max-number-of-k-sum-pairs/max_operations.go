package _679_max_number_of_k_sum_pairs

func maxOperations(nums []int, k int) int {
	mapK := map[int]int{}
	count := 0
	for _, v := range nums {
		if val, ok := mapK[v]; ok && val > 0 {
			mapK[v] = val - 1
			count++
		} else {
			mapK[k-v]++
		}
	}
	return count
}
