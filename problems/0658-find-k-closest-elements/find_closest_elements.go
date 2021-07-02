package _658_find_k_closest_elements

func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k
	for left < right {
		mid := (left + right) / 2
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return arr[left : left+k]
}
