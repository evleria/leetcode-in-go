package _629_slowest_key

func slowestKey(releaseTimes []int, keysPressed string) byte {
	max, result := releaseTimes[0], keysPressed[0]
	for i := 1; i < len(releaseTimes); i++ {
		if r, k := releaseTimes[i]-releaseTimes[i-1], keysPressed[i]; r > max || r == max && k > result {
			max, result = r, k
		}
	}
	return result
}
