package _492_construct_the_rectangle

import "math"

func constructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area)))
	for ; area%w != 0; w-- {
	}
	return []int{area / w, w}
}
