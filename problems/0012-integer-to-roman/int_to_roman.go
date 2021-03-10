package _012_integer_to_roman

func intToRoman(num int) string {
	type node struct {
		value   int
		letters string
	}
	nodes := []node{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result []byte
	for _, n := range nodes {
		for num >= n.value {
			num, result = num-n.value, append(result, []byte(n.letters)...)
		}
	}
	return string(result)
}
