package _775_global_and_local_inversions

func isIdealPermutation(A []int) bool {
	for i, n := range A {
		if d := n - i; d > 1 || d < -1 {
			return false
		}
	}

	return true
}
