package _074_search_a_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	top, bottom := 0, m-1
	for top < bottom {
		mid := (top + bottom) / 2
		switch {
		case matrix[mid][0] == target:
			return true
		case matrix[mid][0] > target:
			bottom = mid - 1
		case matrix[mid][n-1] >= target:
			top, bottom = mid, mid
		default:
			top = mid + 1
		}
	}
	left, right := 0, n-1
	for left < right {
		mid := (left + right) / 2
		switch {
		case matrix[top][mid] == target:
			return true
		case matrix[top][mid] > target:
			right = mid - 1
		default:
			left = mid + 1
		}
	}
	return matrix[top][left] == target
}
