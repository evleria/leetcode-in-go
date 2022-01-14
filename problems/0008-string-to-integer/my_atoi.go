package _008_string_to_integer

func myAtoi(s string) int {
	var rst int
	var sign = 1
	i := 0
	for i < len(s) && ' ' == s[i] {
		i++
	}
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	intMax := 1<<31 - 1
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		if rst > intMax/10 || (rst == intMax/10 && int(s[i]-'0') > intMax%10) {
			if sign == -1 {
				return -1 << 31
			}
			return intMax
		}
		rst = rst*10 + int(s[i]-'0')
		i++
	}
	return rst * sign
}
