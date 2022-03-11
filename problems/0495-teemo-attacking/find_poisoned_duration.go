package _495_teemo_attacking

func findPoisonedDuration(timeSeries []int, duration int) int {
	total := 0
	for i := 0; i < len(timeSeries)-1; i++ {
		diff := timeSeries[i+1] - timeSeries[i]
		if diff > duration {
			diff = duration
		}
		total += diff
	}
	return total + duration
}
