package _850_rectangle_area_II

func rectangleArea(rectangles [][]int) int {
	simpleRects := make([][4]int, 0, len(rectangles))
	for _, cur := range rectangles {
		addRectangle([4]int{cur[0], cur[1], cur[2], cur[3]}, &simpleRects, 0)
	}
	sq := 0
	for _, rectangle := range simpleRects {
		sq += (rectangle[2] - rectangle[0]) * (rectangle[3] - rectangle[1])
	}
	return sq % 1_000_000_007
}

func addRectangle(cur [4]int, rectangles *[][4]int, start int) {
	if start == len(*rectangles) {
		*rectangles = append(*rectangles, cur)
		return
	}

	cpart := (*rectangles)[start]
	if cpart[0] >= cur[2] || cpart[1] >= cur[3] || cpart[2] <= cur[0] || cpart[3] <= cur[1] {
		addRectangle(cur, rectangles, start+1)
		return
	}

	if cur[0] < cpart[0] && cpart[0] < cur[2] {
		addRectangle([4]int{cur[0], cur[1], cpart[0], cur[3]}, rectangles, start+1)
	}
	if cur[0] < cpart[2] && cpart[2] < cur[2] {
		addRectangle([4]int{cpart[2], cur[1], cur[2], cur[3]}, rectangles, start+1)
	}
	if cur[1] < cpart[1] && cpart[1] < cur[3] {
		addRectangle([4]int{cur[0], cur[1], cur[2], cpart[1]}, rectangles, start+1)
	}
	if cur[1] < cpart[3] && cpart[3] < cur[3] {
		addRectangle([4]int{cur[0], cpart[3], cur[2], cur[3]}, rectangles, start+1)
	}
}
