package _943_find_the_shortest_superstring

import (
	"math"
	"strings"
)

func shortestSuperstring(words []string) string {
	n := len(words)
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			graph[i][j] = calculateAdditionalLength(words[i], words[j])
		}
	}

	pow2n := 1 << n
	dp, path := make([][]int, pow2n), make([][]int, pow2n)
	for i := 0; i < pow2n; i++ {
		dp[i], path[i] = make([]int, n), make([]int, n)
	}
	last, min := -1, math.MaxInt64

	for i := 1; i < pow2n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt64
			pow2j := 1 << j
			if i&pow2j > 0 {
				prev := i - pow2j
				if prev == 0 {
					dp[i][j] = len(words[j])
				} else {
					for k := 0; k < n; k++ {
						if dp[prev][k] < math.MaxInt64 && dp[prev][k]+graph[k][j] < dp[i][j] {
							dp[i][j] = dp[prev][k] + graph[k][j]
							path[i][j] = k
						}
					}
				}
			}
			if i == pow2n-1 && dp[i][j] < min {
				min = dp[i][j]
				last = j
			}
		}
	}

	cur := pow2n - 1
	stack := make([]int, 0, n)
	for cur > 0 {
		stack = append(stack, last)
		cur, last = cur-(1<<last), path[cur][last]
	}

	i := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	var sb strings.Builder
	sb.WriteString(words[i])
	for len(stack) > 0 {
		j := stack[len(stack)-1]
		sb.WriteString(words[j][len(words[j])-graph[i][j]:])
		i, stack = j, stack[:len(stack)-1]
	}

	return sb.String()
}

func calculateAdditionalLength(a, b string) int {
	for i := 0; i < len(a); i++ {
		if strings.HasPrefix(b, a[i:]) {
			return len(b) - len(a) + i
		}
	}
	return len(b)
}
