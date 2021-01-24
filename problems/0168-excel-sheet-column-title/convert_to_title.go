package _168_excel_sheet_column_title

func convertToTitle(n int) string {
	var result []byte
	for ; n > 0; n /= 26 {
		n--
		ch := byte(n%26 + 'A')
		result = append([]byte{ch}, result...)
	}
	return string(result)
}
