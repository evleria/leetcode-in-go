package _763_partition_labels

func partitionLabels(s string) []int {
	last := [26]int{}
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}
	j, anchor, result := 0, 0, []int{}
	for i := 0; i < len(s); i++ {
		j = max(j, last[s[i]-'a'])
		if i == j {
			result = append(result, i-anchor+1)
			anchor = i + 1
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
