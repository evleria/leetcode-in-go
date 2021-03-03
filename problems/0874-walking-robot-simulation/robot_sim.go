package _874_walking_robot_simulation

func robotSim(commands []int, obstacles [][]int) int {
	obstaclesMap := make(map[[2]int]bool, len(obstacles))
	for _, obstacle := range obstacles {
		obstaclesMap[[2]int{obstacle[0], obstacle[1]}] = true
	}
	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	cur, d, max := [2]int{}, 0, 0
	for _, cmd := range commands {
		switch cmd {
		case -2:
			d = (d + 3) % 4
		case -1:
			d = (d + 1) % 4
		default:
			curDir := dirs[d]
			for i := 0; i < cmd; i++ {
				pot := [2]int{cur[0] + curDir[0], cur[1] + curDir[1]}
				if obstaclesMap[pot] {
					break
				}
				cur = pot
			}
			if l := cur[0]*cur[0] + cur[1]*cur[1]; l > max {
				max = l
			}
		}
	}
	return max
}
