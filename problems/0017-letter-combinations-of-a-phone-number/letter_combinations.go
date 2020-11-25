package _017_letter_combinations_of_a_phone_number

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	numbers := [][]byte{
		[]byte("abc"),
		[]byte("def"),
		[]byte("ghi"),
		[]byte("jkl"),
		[]byte("mno"),
		[]byte("pqrs"),
		[]byte("tuv"),
		[]byte("wxyz"),
	}

	result := make([]string, 0, 16)
	backtrack(digits, "", &result, numbers)
	return result
}

func backtrack(left string, base string, result *[]string, numbers [][]byte) {
	if len(left) == 0 {
		*result = append(*result, base)
	} else {
		for _, ch := range numbers[left[0]-'2'] {
			backtrack(left[1:], string(append([]byte(base), ch)), result, numbers)
		}
	}
}
