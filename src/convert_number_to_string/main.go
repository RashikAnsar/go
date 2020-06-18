package main

import (
	"fmt"
	"strconv"
)

func main() {
	// numberStr := strconv.FormatInt(100, 10)
	number := 100
	numberStr := strconv.Itoa(number)
	fmt.Println(numberStr)

	numberFloat := 3.14159265358979
	numberFloatStr := strconv.FormatFloat(numberFloat, 'f', 5, 64)
	fmt.Println(numberFloatStr)
	numberFloatStr = strconv.FormatFloat(numberFloat, 'f', 3, 64)
	fmt.Println(numberFloatStr)
}
