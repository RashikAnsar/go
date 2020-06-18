package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	fmt.Println(current.String())
	// https://golang.org/pkg/time/#Parse
	// https://play.golang.org/p/5kULop6obM
	// https://godoc.org/time#example-Time-Format
	fmt.Println("MM-DD-YYYY : ", current.Format("01-02-2006"))
	fmt.Println("YYYY-MM-DD hh:mm:ss : ", current.Format("2006-01-02 15:04:05"))
}
