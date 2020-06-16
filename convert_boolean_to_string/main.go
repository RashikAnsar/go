package main

import (
	"fmt"
)

func main() {
	isNew := true
	// one way to convert using strconv
	// message := "Purchased item is " + strconv.FormatBool(isNew)
	// fmt.Println(message)
	message := "Purchased item is "
	fmt.Printf("%s %v", message, isNew)
}
