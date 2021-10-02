package _174_dungeon_game

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			switch {
			case i == m-1 && j == n-1:
				dungeon[i][j] = max(1, 1-dungeon[i][j])
			case i == m-1:
				dungeon[i][j] = max(1, dungeon[i][j+1]-dungeon[i][j])
			case j == n-1:
				dungeon[i][j] = max(1, dungeon[i+1][j]-dungeon[i][j])
			default:
				dungeon[i][j] = max(1, min(dungeon[i+1][j], dungeon[i][j+1])-dungeon[i][j])
			}
		}
	}

	return dungeon[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
