package _228_summary_ranges

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	start := nums[0]
	intervals := []string{}
	appendInterval := func(start, end int) {
		if start == end {
			intervals = append(intervals, strconv.Itoa(start))
		} else {
			intervals = append(intervals, fmt.Sprintf("%d->%d", start, end))
		}
	}

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != 1 {
			appendInterval(start, nums[i-1])
			start = nums[i]
		}
	}
	appendInterval(start, nums[len(nums)-1])

	return intervals
}
