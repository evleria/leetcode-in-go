package _006_count_number_of_pairs_with_absolute_difference_k

func countKDifference(nums []int, k int) int {
	count := 0
	m := [100]int{}
	for _, num := range nums {
		m[num-1]++
	}
	for i := 0; i < len(m)-k; i++ {
		count += m[i] * m[i+k]
	}
	return count
}
