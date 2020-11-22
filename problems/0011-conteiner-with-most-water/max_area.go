package _011_conteiner_with_most_water

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max, volume := 0, 0

	for left < right {
		dist := right - left
		if hl, hr := height[left], height[right]; hl < hr {
			volume = hl * dist
			left++
		} else {
			volume = hr * dist
			right--
		}

		if volume > max {
			max = volume
		}
	}

	return max
}
