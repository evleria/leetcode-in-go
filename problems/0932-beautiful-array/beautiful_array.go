package _932_beautiful_array

func beautifulArray(n int) []int {
	result := []int{1}
	for len(result) < n {
		temp := make([]int, 0, 2*len(result))
		for _, i := range result {
			if x := i*2 - 1; x <= n {
				temp = append(temp, x)
			}
			if x := i * 2; x <= n {
				temp = append(temp, x)
			}
		}
		result = temp
	}
	return result
}
