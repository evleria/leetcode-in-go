package _641_count_sorted_vowel_strings

func countVowelStrings(n int) int {
	arr := [5]int{1, 1, 1, 1, 1}
	for i := 0; i < n-1; i++ {
		for j := 1; j < 5; j++ {
			arr[j] += arr[j-1]
		}
	}
	count := 0
	for _, num := range arr {
		count += num
	}
	return count
}
