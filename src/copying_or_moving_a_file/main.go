package main

import (
	"io"
	"os"
)

func main() {
	names, err := os.Open("names.txt")

	if err != nil {
		panic(err)
	}

	defer names.Close()

	// namesCopy, err2 := os.Create("target/copy.txt")
	namesCopy, err2 := os.Create("copy.txt")
	if err2 != nil {
		panic(err2)
	}

	defer namesCopy.Close()
	_, err3 := io.Copy(namesCopy, names)

	if err3 != nil {
		panic(err3)
	}
}
