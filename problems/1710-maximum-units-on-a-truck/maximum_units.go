package _710_maximum_units_on_a_truck

import (
	"math"
	"sort"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	sum := 0
	for i := 0; i < len(boxTypes) && truckSize > 0; i++ {
		boxes := int(math.Min(float64(boxTypes[i][0]), float64(truckSize)))
		sum, truckSize = sum+boxes*boxTypes[i][1], truckSize-boxes
	}
	return sum
}
