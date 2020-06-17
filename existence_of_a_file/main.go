package main

import (
	"fmt"
	"os"
)

func main() {
	// if _, err := os.Stat("log.txt"); err == nil {
	// 	fmt.Println("log.txt file exists")
	// } else {
	// 	fmt.Println("No such file")
	// }

	if _, err := os.Stat("log.txt"); os.IsNotExist(err) {
		fmt.Println("log.txt file does not exists")
	} else {
		fmt.Println("log.txt file exists")
	}
}
