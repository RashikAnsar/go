package main

import "fmt"

func main() {
	numbers := []int{1, 3, 5, 9, 7, 6, 2, 10, 8, 4}

	// for index and value
	for idx, value := range numbers {
		fmt.Printf("Index: %v\tValue: %v\n", idx, value)
	}

	// for value
	for _, val := range numbers {
		fmt.Printf("%v\t", val)
	}
}
