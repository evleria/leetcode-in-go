package _852_peak_index_in_a_mountain_array

func peakIndexInMountainArray(arr []int) int {
	i := 0
	for arr[i] < arr[i+1] {
		i++
	}
	return i
}
