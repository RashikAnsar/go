package main

func isPowerOfTwo(x int) bool {
	// if x < 1 {
	// 	return false
	// }
	// for x != 1 {
	// 	if x%2 != 0 {
	// 		return false
	// 	}
	// 	x = x / 2
	// }
	// return true
	// // using bitwise operations
	return x&(x-1) == 0
}

func main() {}
