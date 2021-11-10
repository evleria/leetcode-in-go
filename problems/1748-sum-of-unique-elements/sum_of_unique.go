package _748_sum_of_unique_elements

func sumOfUnique(nums []int) int {
	sum := 0
	m := [100]int{}
	for _, num := range nums {
		m[num-1]++
	}
	for i, n := range m {
		if n == 1 {
			sum += i + 1
		}
	}
	return sum
}
