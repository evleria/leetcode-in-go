package _282_expression_add_operators

import "strconv"

func addOperators(num string, target int) []string {
	result := []string{}

	backtrack(&result, num, "", target, 0, 0)

	return result
}

func backtrack(result *[]string, num, s string, target, val, last int) {
	if len(num) == 0 {
		if val == target {
			*result = append(*result, s)
		}
		return
	}

	for i := 1; i <= len(num); i++ {
		str := num[:i]
		digit, _ := strconv.Atoi(str)

		if num[0] == '0' && i != 1 {
			continue
		}

		if s == "" {
			backtrack(result, num[i:], str, target, digit, digit)
		} else {
			backtrack(result, num[i:], s+"+"+str, target, val+digit, digit)
			backtrack(result, num[i:], s+"-"+str, target, val-digit, -digit)
			backtrack(result, num[i:], s+"*"+str, target, val-last+last*digit, last*digit)
		}
	}
}
