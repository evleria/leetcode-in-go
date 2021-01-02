package _006_zigzag_conversion

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	result, cycle := make([]byte, 0, len(s)), numRows*2-2
	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += cycle {
			result = append(result, s[j])
			if x := j + cycle - 2*i; i != 0 && i != numRows-1 && x < len(s) {
				result = append(result, s[x])
			}
		}
	}

	return string(result)
}
