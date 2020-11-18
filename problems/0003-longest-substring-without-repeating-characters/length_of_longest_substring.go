package _003_longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	max, start := 0, 0

	dict := map[rune]int{}

	for i, ch := range s {
		if j, ok := dict[ch]; ok && j >= start {
			start = j + 1
		}

		dict[ch] = i
		if l := i - start + 1; l > max {
			max = l
		}
	}

	return max
}
