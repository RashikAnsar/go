package main

import (
	"fmt"
	"time"
)

func main() {

	// main application also runs under goroutine
	// goroutine give control back to other goroutines voluntarily instead
	// of premptively

	// There are two types of channels
	// 1. Buffered
	// 2. UnBuffered

	nameChannel := make(chan string)

	go func() {
		names := []string{"Michael", "Franklin", "Trevor", "Lester", "Benny"}

		for _, name := range names {
			time.Sleep(1 * time.Second)
			// fmt.Println(name)
			nameChannel <- name
		}
		close(nameChannel)
	}()

	// fmt.Println(<-nameChannel)
	for data := range nameChannel {
		fmt.Println(data)
	}

	// go func() {
	// 	ages := []int{34, 28, 43, 32, 30}

	// 	for _, age := range ages {
	// 		time.Sleep(1 * time.Second)
	// 		fmt.Println(age)
	// 	}
	// }()

	// time.Sleep(10 * time.Second)
}
