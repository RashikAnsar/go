package main

import (
	"errors"
	"fmt"
)

func main() {
	// parsedDate, err := time.Parse("2006", "2020")
	// if err != nil {
	// 	fmt.Println("An error occured ::", err.Error())
	// } else {
	// 	fmt.Println(parsedDate)
	// }

	_, err := doSomething()
	if err != nil {
		fmt.Println(err)
	}
}

func doSomething() (string, error) {
	return "somethin", errors.New("something went wrong")
}
