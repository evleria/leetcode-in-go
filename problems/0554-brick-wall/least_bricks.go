package _554_brick_wall

func leastBricks(wall [][]int) int {
	m := map[int]int{}
	max := 0
	for _, row := range wall {
		offset := 0
		for i := 0; i < len(row)-1; i++ {
			offset += row[i]
			v := m[offset] + 1
			if v > max {
				max = v
			}
			m[offset] = v
		}
	}

	return len(wall) - max
}
