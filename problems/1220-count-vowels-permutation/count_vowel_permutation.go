package _220_count_vowels_permutation

const mod = 1_000_000_007

func countVowelPermutation(n int) int {
	a, e, i, o, u := 1, 1, 1, 1, 1
	for j := 0; j < n-1; j++ {
		a, e, i, o, u = (e+i+u)%mod, (a+i)%mod, (e+o)%mod, i, (o+i)%mod
	}
	return (a + e + i + o + u) % mod
}
