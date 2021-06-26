package _688_count_of_matches_in_tournament

func numberOfMatches(n int) int {
	matches := 0
	for n > 1 {
		lucky := n % 2
		n, matches = lucky+n/2, matches+n/2
	}

	return matches
}
