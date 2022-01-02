package _010_pairs_of_songs_with_total_durations_divisible_by_60

func numPairsDivisibleBy60(time []int) int {
	count, complements := 0, [60]int{}
	for _, t := range time {
		count += complements[(600-t)%60]
		complements[t%60]++
	}

	return count
}
