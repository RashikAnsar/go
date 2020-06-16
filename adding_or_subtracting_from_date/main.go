package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	fmt.Println(current.String())

	tenMoreMinutes := current.Add(10 * time.Minute)
	fmt.Println(tenMoreMinutes)

	nextMonthDate := current.AddDate(0, 1, 0)
	fmt.Println(nextMonthDate.String())

	oneLessYears := nextMonthDate.AddDate(-1, 0, 0)
	fmt.Println(oneLessYears.String())

}
