package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(readLine(3))
	fmt.Println(readLine(12))
	fmt.Println(readLine(18))
}

func readLine(lineNumber int) string {
	file, err := os.Open("names.txt")

	defer file.Close()

	if err != nil {
		log.Fatalln("ERROR: ", err)
	}

	fileScanner := bufio.NewScanner(file)
	lineCount := 0

	for fileScanner.Scan() {
		if lineCount == lineNumber {
			return fileScanner.Text()
		}
		lineCount++
	}
	return "THIS LINE IS EMPTY (DOES NOT EXIST)"
}
