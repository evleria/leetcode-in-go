package _917_reverse_only_letters

import "unicode"

func reverseOnlyLetters(s string) string {
	str := []rune(s)
	left, right := 0, len(str)-1
	for left < right {
		for ; left < right && !unicode.IsLetter(str[left]); left++ {
		}
		for ; left < right && !unicode.IsLetter(str[right]); right-- {
		}
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}

	return string(str)
}
