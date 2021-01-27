package _412_fizz_buzz

import "strconv"

func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			result[i-1] = "FizzBuzz"
		case i%3 == 0:
			result[i-1] = "Fizz"
		case i%5 == 0:
			result[i-1] = "Buzz"

		default:
			result[i-1] = strconv.Itoa(i)
		}
	}

	return result
}
