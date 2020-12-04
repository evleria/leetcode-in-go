package _171_excel_sheet_column_number

func titleToNumber(s string) int {
	sum := 0
	for _, ch := range s {
		sum *= 26
		sum += int(ch-'A') + 1
	}

	return sum
}
