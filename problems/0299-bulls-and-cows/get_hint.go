package _299_bulls_and_cows

import "fmt"

func getHint(secret string, guess string) string {
	bulls, cows := 0, 0
	nums := [10]int{}
	for i := 0; i < len(secret); i++ {
		if s, g := secret[i]-'0', guess[i]-'0'; s == g {
			bulls++
		} else {
			if nums[s] < 0 {
				cows++
			}
			if nums[g] > 0 {
				cows++
			}
			nums[s]++
			nums[g]--
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}
