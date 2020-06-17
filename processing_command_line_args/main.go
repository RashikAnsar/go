package main

import (
	"fmt"
	"os"
)

func main() {
	realArgs := os.Args[1:]
	if len(realArgs) == 0 {
		fmt.Println("No arguments passed")
		return
	}

	switch realArgs[0] {
	case "a":
		writeHelloWorld()
	case "b":
		writeHelloMars()
	default:
		fmt.Println("Please pass a valid argument 'a', 'b'")
	}
}

func writeHelloWorld() {
	fmt.Println("Hello, World!")
}

func writeHelloMars() {
	fmt.Println("Hello, Mars...")
}
