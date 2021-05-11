package _423_maximum_points_you_can_obtain_from_gards

func maxScore(cardPoints []int, k int) int {
	sum := 0
	for i := 0; i < k; i++ {
		sum += cardPoints[i]
	}
	if len(cardPoints) == k {
		return sum
	}
	max, left, right := sum, k-1, len(cardPoints)-1
	for i := 0; i < k; i++ {
		sum += cardPoints[right-i] - cardPoints[left-i]
		if sum > max {
			max = sum
		}
	}

	return max
}
