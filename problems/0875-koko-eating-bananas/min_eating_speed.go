package _875_koko_eating_bananas

import "math"

func minEatingSpeed(piles []int, h int) int {
	max := piles[0]
	for i := 1; i < len(piles); i++ {
		if piles[i] > max {
			max = piles[i]
		}
	}
	ratio := h / len(piles)
	max = int(math.Ceil(float64(max) / float64(ratio)))

	min := 1
	for min < max {
		mid := (min + max) / 2
		if enoughTime(piles, mid, h) {
			max = mid
		} else {
			min = mid + 1
		}
	}

	return min
}

func enoughTime(piles []int, k, h int) bool {
	for _, pile := range piles {
		h -= int(math.Ceil(float64(pile) / float64(k)))
		if h < 0 {
			return false
		}
	}
	return true
}
