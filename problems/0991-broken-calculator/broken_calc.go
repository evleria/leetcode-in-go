package _991_broken_calculator

func brokenCalc(startValue int, target int) int {
	steps := 0
	for target > startValue {
		if target%2 == 0 {
			target /= 2
		} else {
			target++
		}
		steps++
	}
	return steps + startValue - target
}
