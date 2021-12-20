package _200_minimum_absolute_difference

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	result, diff := [][]int{}, math.MaxInt64
	for i := 1; i < len(arr); i++ {
		nDiff := arr[i] - arr[i-1]
		if nDiff == diff {
			result = append(result, []int{arr[i-1], arr[i]})
		} else if nDiff < diff {
			result, diff = [][]int{{arr[i-1], arr[i]}}, arr[i]-arr[i-1]
		}
	}

	return result
}
