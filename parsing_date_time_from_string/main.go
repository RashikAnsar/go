package main

import (
	"fmt"
	"time"
)

func main() {
	str := "2020-06-12T11:45:26.371Z"
	layout := "2006-01-02T15:04:05.000Z"

	t, err := time.Parse(layout, str)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(t.String())
}
