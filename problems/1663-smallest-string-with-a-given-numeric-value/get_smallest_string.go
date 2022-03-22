package _663_smallest_string_with_a_given_numeric_value

func getSmallestString(n int, k int) string {
	result := make([]rune, n)
	d := k - n
	for i := n - 1; i >= 0; i-- {
		delta := d
		if delta > 25 {
			delta = 25
		}

		result[i] = rune('a' + delta)
		d -= delta
	}

	return string(result)
}
