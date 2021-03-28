package _423_reconstruct_original_digits_from_english

import (
	"strconv"
	"strings"
)

var digits = [10]struct {
	digit int
	char  byte
	word  string
}{
	{0, 'z', "zero"},
	{2, 'w', "two"},
	{4, 'u', "four"},
	{6, 'x', "six"},
	{8, 'g', "eight"},
	{1, 'o', "one"},
	{3, 'r', "three"},
	{5, 'f', "five"},
	{7, 'v', "seven"},
	{9, 'i', "nine"},
}

func originalDigits(s string) string {
	counts := [10]int{}
	chars := convert(s)
	for _, digit := range digits {
		n := chars[digit.char-'a']
		counts[digit.digit] = n
		wordChars := convert(digit.word)
		for i := 0; i < 26; i++ {
			chars[i] -= n * wordChars[i]
		}
	}

	var result strings.Builder
	for i := 0; i < 10; i++ {
		result.WriteString(strings.Repeat(strconv.Itoa(i), counts[i]))
	}
	return result.String()
}

func convert(s string) [26]int {
	chars := [26]int{}
	for i := 0; i < len(s); i++ {
		chars[s[i]-'a']++
	}
	return chars
}
