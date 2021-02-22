package _389_find_the_difference

func findTheDifference(s string, t string) byte {
	var x byte
	for i := 0; i < len(s); i++ {
		x ^= s[i] ^ t[i]
	}
	return x ^ t[len(t)-1]
}
