package _465_maximum_area_of_a_piece_of_cake_after_horizontal_and_vertical_cuts

import "sort"

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	result := maxGap(horizontalCuts, h) * maxGap(verticalCuts, w)

	return result % 1_000_000_007
}

func maxGap(nums []int, end int) int {
	max, cur := end-nums[len(nums)-1], 0
	for _, n := range nums {
		if gap := n - cur; gap > max {
			max = gap
		}
		cur = n
	}

	return max
}
