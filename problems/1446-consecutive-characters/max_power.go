package _446_consecutive_characters

func maxPower(s string) int {
	count, maxCount := 1, 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			count = 1
		}
	}
	return maxCount
}
