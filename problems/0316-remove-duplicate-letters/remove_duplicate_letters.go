package _316_remove_duplicate_letters

func removeDuplicateLetters(s string) string {
	count := map[byte]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}
	stack, visited := []byte{}, map[byte]bool{}
	for i := 0; i < len(s); i++ {
		count[s[i]]--
		if visited[s[i]] {
			continue
		}
		for len(stack) != 0 && s[i] < stack[len(stack)-1] && count[stack[len(stack)-1]] != 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			visited[top] = false
		}
		stack = append(stack, s[i])
		visited[s[i]] = true
	}
	return string(stack)
}
