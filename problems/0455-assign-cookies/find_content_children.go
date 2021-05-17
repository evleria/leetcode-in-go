package _455_assign_cookies

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var i int
	for j := 0; i < len(g) && j < len(s); j++ {
		if g[i] <= s[j] {
			i++
		}
	}
	return i
}
