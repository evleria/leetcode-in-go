package _067_add_binary

func addBinary(a string, b string) string {
	var carry byte
	if len(a) < len(b) {
		a, b = b, a
	}

	result := make([]byte, len(a)+1)
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || carry != 0; {
		sum := carry
		if i >= 0 {
			sum += a[i] - '0'
		}
		if j >= 0 {
			sum += b[j] - '0'
		}
		result[i+1] = sum%2 + '0'
		carry = sum / 2
		i--
		j--
	}

	if result[0] == '0' {
		result = result[1:]
	}
	return string(result)
}
