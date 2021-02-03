package _351_count_negative_numbers_in_a_sorted_matrix

func countNegatives(grid [][]int) int {
	acc := 0
	for _, row := range grid {
		acc += countRow(row)
	}
	return acc
}

func countRow(row []int) int {
	if row[0] < 0 {
		return len(row)
	}

	left, right := 0, len(row)-1
	for left <= right {
		mid := (left + right) / 2
		if row[mid] >= 0 {
			left = mid + 1
		} else {
			if mid != 0 && row[mid-1] < 0 {
				right = mid - 1
			} else {
				return len(row) - mid
			}
		}
	}

	return 0
}
