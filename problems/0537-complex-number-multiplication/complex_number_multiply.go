package _537_complex_number_multiplication

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {
	r1, i1 := parseComplexNumber(num1)
	r2, i2 := parseComplexNumber(num2)
	return fmt.Sprintf("%d+%di", r1*r2-i1*i2, r1*i2+r2*i1)
}

func parseComplexNumber(num string) (int, int) {
	idx := strings.IndexByte(num, '+')
	r, _ := strconv.Atoi(num[:idx])
	i, _ := strconv.Atoi(num[idx+1 : len(num)-1])
	return r, i
}
