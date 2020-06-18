package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := "2"
	// strconv.Atoi()
	valueInt, err := strconv.ParseInt(number, 10, 32)
	if err != nil {
		fmt.Println("Parsing from number to string failed")
	} else {
		if valueInt == 2 {
			fmt.Println("Success")
		} else {
			fmt.Println(valueInt)
		}
	}

	numberFloat := "3.14"
	valueFloat, errFloat := strconv.ParseFloat(numberFloat, 64)
	if errFloat != nil {
		fmt.Println("Parsing from number to string failed")
	} else {
		if valueFloat == 3.14 {
			fmt.Println("Success")
		} else {
			fmt.Println(valueFloat)
		}
	}
}
