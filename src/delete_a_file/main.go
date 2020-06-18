package main

import "os"

func main() {
	err := os.Remove("names.txt")

	if err != nil {
		panic(err)
	}
}
