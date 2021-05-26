package _689_partitioning_into_minimum_number_of_deci_binary_numbers

func minPartitions(n string) int {
	var max byte = '0'
	for i := 0; i < len(n); i++ {
		if n[i] > max {
			max = n[i]
		}
	}
	return int(max - '0')
}
