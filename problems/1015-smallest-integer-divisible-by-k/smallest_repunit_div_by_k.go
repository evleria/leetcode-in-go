package _015_smallest_integer_divisible_by_k

func smallestRepunitDivByK(k int) int {
	r := 0
	for l := 1; l <= k; l++ {
		r = (r*10 + 1) % k
		if r == 0 {
			return l
		}
	}
	return -1
}
