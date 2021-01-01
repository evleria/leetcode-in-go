package _056_merge_intervals

import "sort"

func merge(intervals [][]int) [][]int {
	var result [][]int

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	for i := 0; i < len(intervals); i++ {
		if i == len(intervals)-1 || intervals[i][1] < intervals[i+1][0] {
			result = append(result, intervals[i])
		} else {
			start, end := intervals[i][0], intervals[i][1]
			for j := i + 1; j < len(intervals) && intervals[j][0] <= end; j++ {
				if x := intervals[j][1]; x > end {
					end = x
				}
				i++
			}
			result = append(result, []int{start, end})
		}
	}
	return result
}
