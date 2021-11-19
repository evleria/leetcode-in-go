package _657_robot_return_to_origin

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, m := range moves {
		switch m {
		case 'R':
			x++
		case 'L':
			x--
		case 'U':
			y++
		case 'D':
			y--
		}
	}
	return x == 0 && y == 0
}
