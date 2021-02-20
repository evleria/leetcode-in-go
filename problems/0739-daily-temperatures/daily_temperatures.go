package _739_daily_temperatures

func dailyTemperatures(T []int) []int {
	type Day struct {
		Temp  int
		Index int
	}
	result := make([]int, len(T))
	var stack []Day
	for i := len(T) - 1; i >= 0; i-- {
		temp := T[i]
		for len(stack) > 0 && temp >= stack[len(stack)-1].Temp {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			result[i] = stack[len(stack)-1].Index - i
		}
		stack = append(stack, Day{Index: i, Temp: temp})
	}
	return result
}
