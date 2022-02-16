package _725_number_of_rectangles_that_can_form_the_largest_square

func countGoodRectangles(rectangles [][]int) int {
	count, maxSize := 0, 0
	for _, rect := range rectangles {
		size := min(rect[0], rect[1])
		switch {
		case size > maxSize:
			count, maxSize = 1, size
		case size == maxSize:
			count++
		}
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
