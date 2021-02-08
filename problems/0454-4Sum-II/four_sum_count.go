package _454_4Sum_II

func fourSumCount(A []int, B []int, C []int, D []int) int {
	m := map[int]int{}
	for _, a := range A {
		for _, b := range B {
			m[a+b]++
		}
	}

	counter := 0
	for _, c := range C {
		for _, d := range D {
			counter += m[-(c + d)]
		}
	}
	return counter
}
