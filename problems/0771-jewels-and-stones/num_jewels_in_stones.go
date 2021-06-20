package _771_jewels_and_stones

func numJewelsInStones(jewels string, stones string) int {
	isJewel := make(map[rune]bool, len(jewels))
	for _, r := range jewels {
		isJewel[r] = true
	}

	count := 0
	for _, s := range stones {
		if isJewel[s] {
			count++
		}
	}
	return count
}
