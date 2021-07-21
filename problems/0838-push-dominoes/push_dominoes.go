package _838_push_dominoes

func pushDominoes(dominoes string) string {
	result := []byte(dominoes)

	setInterval := func(from, to int, dir byte) {
		for i := from; i < to; i++ {
			result[i] = dir
		}
	}

	lastR, lastAffected := -1, -1
	for i, d := range result {
		switch d {
		case 'L':
			if lastR != -1 {
				for left, right := lastR+1, i-1; left < right; left, right = left+1, right-1 {
					result[left], result[right] = 'R', 'L'
				}
				lastR = -1
			} else {
				setInterval(lastAffected+1, i, 'L')
			}
			lastAffected = i
		case 'R':
			if lastR != -1 {
				setInterval(lastR+1, i, 'R')
			}
			lastR, lastAffected = i, i
		}
	}
	if lastR != -1 {
		setInterval(lastR, len(result), 'R')
	}

	return string(result)
}
