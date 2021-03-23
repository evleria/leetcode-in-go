package _923_3Sum_with_multiplicity

func threeSumMulti(arr []int, target int) int {
	occurrences, sums, result := make([]int, 101), make([]int, 201), 0
	for _, n := range arr {
		if idx := target - n; idx >= 0 && idx < len(sums) {
			result += sums[idx]
		}
		for i, o := range occurrences {
			sums[i+n] += o
		}
		occurrences[n]++
	}
	return result % 1_000_000_007
}
