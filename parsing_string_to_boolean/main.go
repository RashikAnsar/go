package main

import (
	"fmt"
	"strconv"
)

func main() {
	// try with isNew = "false" also
	isNew := "true"
	isNewBool, err := strconv.ParseBool(isNew)
	if err != nil {
		fmt.Println("Failed to parse boolean as string")
	} else {
		if isNewBool {
			fmt.Println("Is New")
		} else {
			fmt.Println("Not new")
		}
	}

}
