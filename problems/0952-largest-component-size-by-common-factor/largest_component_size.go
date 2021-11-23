package _952_largest_component_size_by_common_factor

import "math"

type DSU struct {
	parent []int
}

func (d *DSU) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) union(x, y int) {
	d.parent[d.find(x)] = d.parent[d.find(y)]
}

func Constructor(n int) *DSU {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}
	return &DSU{s}
}

func largestComponentSize(A []int) int {
	ans, max, cache := 1, 0, make(map[int]int)
	for _, a := range A {
		if a > max {
			max = a
		}
	}
	DSU := Constructor(max + 1)
	for _, a := range A {
		for k := 2; k <= int(math.Sqrt(float64(a))); k++ {
			if a%k == 0 {
				DSU.union(a, k)
				DSU.union(a, a/k)
			}
		}
	}
	for _, a := range A {
		r := DSU.find(a)
		cache[r]++
		if cache[r] > ans {
			ans = cache[r]
		}
	}
	return ans
}
