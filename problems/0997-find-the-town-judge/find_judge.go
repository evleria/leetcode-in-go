package _997_find_the_town_judge

func findJudge(n int, trust [][]int) int {
	tr := make([][2]int, n+1)
	for _, t := range trust {
		who, whom := t[0], t[1]
		tr[who][0]++
		tr[whom][1]++
	}

	for i := 1; i < len(tr); i++ {
		if p := tr[i]; p[0] == 0 && p[1] == n-1 {
			return i
		}
	}
	return -1
}
