package _452_minimum_number_of_arrows_to_burst_balloons

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	arrows := 0
	for i := 0; i < len(points); i++ {
		arrows++
		shotAt := points[i][1]
		for ; i+1 < len(points) && points[i+1][0] <= shotAt; i++ {
		}
	}

	return arrows
}
