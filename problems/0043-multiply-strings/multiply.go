package _043_multiply_strings

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	r := make([]byte, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			r[i+j+1] += (num1[i] - '0') * (num2[j] - '0')
			if r[i+j+1] > 9 {
				r[i+j] += r[i+j+1] / 10
				r[i+j+1] %= 10
			}
		}
	}
	for i := len(r) - 1; i >= 0; i-- {
		r[i] += '0'
	}
	if r[0] == '0' {
		r = r[1:]
	}
	return string(r)
}
