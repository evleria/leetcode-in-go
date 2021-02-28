package _664_ways_to_make_a_fair_array

func waysToMakeFair(nums []int) int {
	result, left, right := 0, [2]int{}, [2]int{}
	for i, v := range nums {
		right[i%2] += v
	}
	for i, v := range nums {
		right[i%2] -= v
		if left[0]+right[1] == left[1]+right[0] {
			result++
		}
		left[i%2] += v
	}
	return result
}
