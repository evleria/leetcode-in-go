package _096_unique_binary_search_trees

func numTrees(n int) int {
	answers := make([]int, n+1)
	answers[0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			answers[i] += answers[j-1] * answers[i-j]
		}
	}

	return answers[n]
}
