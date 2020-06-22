package main

import "fmt"

// compute exponents using (+, -, *, /, %)
func exponent(base int, power int) int {
	result := 1
	for power != 0 {
		if power%2 == 0 {
			base *= base
			power /= 2
		} else {
			result = result * base
			power = power - 1
		}
	}
	return result
}

func main() {
	fmt.Println(exponent(2, 5))
}
