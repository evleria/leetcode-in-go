package _645__set_mismatch

func findErrorNums(nums []int) []int {
	var doubled int
	s := make([]bool, len(nums))
	for _, n := range nums {
		if s[n-1] {
			doubled = n
		}
		s[n-1] = true
	}
	for i, b := range s {
		if !b {
			return []int{doubled, i + 1}
		}
	}
	return nil
}
