package main

// GCD Greatest Common Divisor (Euclidean Algorithm)
func GCD(x int, y int) int {
	// // using loops
	// for x != y {
	// 	if x > y {
	// 		x = x - y
	// 	} else {
	// 		y = y - x
	// 	}
	// }
	// return y

	// // using recursion
	if y == 0 {
		return x
	}

	return GCD(y, x%y)
}

func main() {}
