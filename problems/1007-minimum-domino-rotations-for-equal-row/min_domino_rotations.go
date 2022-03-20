package _007_minimum_domino_rotations_for_equal_row

func minDominoRotations(tops []int, bottoms []int) int {
	var ls [7][3]int
	for i, t := range tops {
		b := bottoms[i]
		ls[t][0]++
		ls[b][1]++
		if t == b {
			ls[t][2]++
		}
	}
	for _, n := range ls {
		if n[0]+n[1]-n[2] == len(tops) {
			if n[0] > n[1] {
				return len(tops) - n[0]
			}
			return len(tops) - n[1]
		}
	}
	return -1
}
