package _415_add_strings

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	result := make([]byte, 0, len(num1)+1)
	var carry byte
	for i, j := len(num1)-1, len(num2)-1; i >= 0; j, i = j-1, i-1 {
		sum := carry + num1[i] - '0'
		if j >= 0 {
			sum += num2[j] - '0'
		}
		carry, sum = sum/10, sum%10
		result = append(result, sum+'0')
	}
	if carry != 0 {
		result = append(result, carry+'0')
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
