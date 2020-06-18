package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data := "hello world"
	err := ioutil.WriteFile("output.txt", []byte(data), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
