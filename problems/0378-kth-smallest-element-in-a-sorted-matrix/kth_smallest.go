package _378_kth_smallest_element_in_a_sorted_matrix

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	start, end, mid := matrix[0][0], matrix[n-1][n-1], 0
	count := 0
	for start < end {
		mid = start + (end-start)/2
		count = countLessEqual(matrix, mid)
		if count < k {
			start = mid + 1
			continue
		}

		end = mid
	}
	return start
}

func countLessEqual(matrix [][]int, mid int) int {
	count := 0
	n := len(matrix)
	for r, c := n-1, 0; r >= 0 && c < n; {
		if matrix[r][c] > mid {
			r--
			continue
		}

		count += r + 1
		c++
	}

	return count
}
