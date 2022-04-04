package _030_matrix_cells_in_distance_order

import "sort"

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	result := make([][]int, 0, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result = append(result, []int{i, j})
		}
	}
	sort.Slice(result, func(i, j int) bool {
		ri, rj := result[i][0], result[i][1]
		rk, rl := result[j][0], result[j][1]
		return abs(ri-rCenter)+abs(rj-cCenter) < abs(rk-rCenter)+abs(rl-cCenter)
	})
	return result
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
