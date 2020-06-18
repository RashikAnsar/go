package main

import "os"

func main() {
	// only directory
	err := os.Remove("hello")
	// // Directory along with sub files/folders
	// err := os.RemoveAll("hello")
	if err != nil {
		panic(err)
	}
}
