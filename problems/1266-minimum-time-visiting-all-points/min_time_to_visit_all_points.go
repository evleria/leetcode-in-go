package _266_minimum_time_visiting_all_points

func minTimeToVisitAllPoints(points [][]int) int {
	count := 0
	for i := 0; i < len(points)-1; i++ {
		dx, dy := abs(points[i+1][0]-points[i][0]), abs(points[i+1][1]-points[i][1])
		count += max(dx, dy)
	}
	return count
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
