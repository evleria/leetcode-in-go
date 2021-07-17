package _927_three_equal_parts

var Invalid = []int{-1, -1}

func threeEqualParts(arr []int) []int {
	ones := make([]int, 0, len(arr))
	for i, n := range arr {
		if n == 1 {
			ones = append(ones, i)
		}
	}

	if len(ones) == 0 {
		return []int{0, len(arr) - 1}
	}
	if len(ones)%3 != 0 {
		return Invalid
	}

	gr := len(ones) / 3
	start1, start2, start3 := ones[0], ones[gr], ones[2*gr]
	grLen := len(arr) - start3

	first, second, third := arr[start1:start1+grLen], arr[start2:start2+grLen], arr[start3:]
	if !equal(first, second, third) {
		return Invalid
	}

	return []int{start1 + grLen - 1, start2 + grLen}
}

func equal(a, b, c []int) bool {
	if len(a) != len(b) || len(a) != len(c) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] || a[i] != c[i] {
			return false
		}
	}

	return true
}
