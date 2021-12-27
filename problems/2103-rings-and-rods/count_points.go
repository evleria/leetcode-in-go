package _103_rings_and_rods

func countPoints(rings string) int {
	rods := [10]int{}
	for i := 0; i < len(rings); i += 2 {
		color, idx := rings[i], rings[i+1]-'0'

		switch color {
		case 'R':
			rods[idx] |= 1 << 0
		case 'G':
			rods[idx] |= 1 << 1
		case 'B':
			rods[idx] |= 1 << 2
		}
	}

	count := 0
	for _, rod := range rods {
		if rod == 0b111 {
			count++
		}
	}
	return count
}
