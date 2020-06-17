package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("names.txt")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	defer file.Close()
	fmt.Println(lineCount)
}
