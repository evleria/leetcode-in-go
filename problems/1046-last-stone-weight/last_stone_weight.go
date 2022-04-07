package _046_last_stone_weight

import "sort"

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}
	sort.Ints(stones)
	for len(stones) >= 2 {
		stone := stones[len(stones)-1] - stones[len(stones)-2]
		stones = stones[:len(stones)-2]
		if stone > 0 {
			idx := sort.SearchInts(stones, stone)
			if idx >= len(stones) {
				stones = append(stones, stone)
			} else {
				stones = append(stones[:idx], append([]int{stone}, stones[idx:]...)...)
			}
		}
	}
	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}
