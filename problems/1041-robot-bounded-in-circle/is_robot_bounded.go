package _041_robot_bounded_in_circle

func isRobotBounded(instructions string) bool {
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	i, x, y := 0, 0, 0
	for j := 0; j < len(instructions); j++ {
		switch instructions[j] {
		case 'R':
			i = (i + 1) % 4
		case 'L':
			i = (i + 3) % 4
		case 'G':
			x += dir[i][0]
			y += dir[i][1]
		}
	}
	return x == 0 && y == 0 || i > 0
}
