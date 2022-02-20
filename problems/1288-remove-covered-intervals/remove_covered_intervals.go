package _288_remove_covered_intervals

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] > intervals[j][1]
	})

	count := len(intervals)
	for i := 0; i < len(intervals)-1; i++ {
		j := i + 1
		for ; j < len(intervals) && intervals[i][1] >= intervals[j][1]; j++ {
			count--
		}
		i = j - 1
	}
	return count
}
