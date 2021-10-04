package _323_maximum_69_number

import "strconv"

func maximum69Number(num int) int {
	str := strconv.Itoa(num)
	m := []byte(str)
	for i, ch := range m {
		if ch == '6' {
			m[i] = '9'
			break
		}
	}
	result, _ := strconv.Atoi(string(m))
	return result
}
