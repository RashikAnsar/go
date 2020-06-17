package main

import "os"

func main() {
	err := os.Rename("old.txt", "new.txt")

	if err != nil {
		panic(err)
	}
}
