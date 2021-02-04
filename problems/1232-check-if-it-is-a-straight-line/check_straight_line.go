package _232_check_if_it_is_a_straight_line

func checkStraightLine(coordinates [][]int) bool {
	first, rest, last := coordinates[0], coordinates[1:len(coordinates)-1], coordinates[len(coordinates)-1]
	dx, dy := last[0]-first[0], last[1]-first[1]
	for _, point := range rest {
		if dx*(point[1]-first[1]) != dy*(point[0]-first[0]) {
			return false
		}
	}

	return true
}
