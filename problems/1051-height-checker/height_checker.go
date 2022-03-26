package _051_height_checker

import "sort"

func heightChecker(heights []int) int {
	count, expected := 0, make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)
	for i := 0; i < len(heights); i++ {
		if heights[i] != expected[i] {
			count++
		}
	}
	return count
}
