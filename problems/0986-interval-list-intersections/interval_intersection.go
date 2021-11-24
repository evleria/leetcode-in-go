package _986_interval_list_intersections

import "math"

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	ans := [][]int{}
	i, j := 0, 0
	for i < len(firstList) && j < len(secondList) {
		lo := int(math.Max(float64(firstList[i][0]), float64(secondList[j][0])))
		hi := int(math.Min(float64(firstList[i][1]), float64(secondList[j][1])))
		if lo <= hi {
			ans = append(ans, []int{lo, hi})
		}
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return ans
}
