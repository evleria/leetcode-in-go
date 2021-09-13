package _189_maximum_number_of_balloons

import "math"

func maxNumberOfBalloons(text string) int {
	m := [26]int{}
	for i := range text {
		m[text[i]-'a']++
	}

	m['l'-'a'] /= 2
	m['o'-'a'] /= 2

	min := math.MaxInt32
	for _, ch := range "balon" {
		if s := m[ch-'a']; s < min {
			min = s
		}
	}

	return min
}
