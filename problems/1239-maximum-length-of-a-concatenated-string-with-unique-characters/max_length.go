package _239_maximum_length_of_a_concatenated_string_with_unique_characters

func maxLength(arr []string) int {
	max := 0
	backtrack(arr, 0, &max, "")
	return max
}

func backtrack(arr []string, i int, max *int, s string) {
	if i == len(arr) && unique(s) {
		if len(s) > *max {
			*max = len(s)
		}
	}
	if i == len(arr) {
		return
	}
	if !unique(s) {
		return
	}

	backtrack(arr, i+1, max, s+arr[i])
	backtrack(arr, i+1, max, s)
}

func unique(s string) bool {
	arr := make([]int, 26)
	for _, v := range s {
		if arr[v-'a'] > 0 {
			return false
		}
		arr[v-'a']++
	}
	return true
}
