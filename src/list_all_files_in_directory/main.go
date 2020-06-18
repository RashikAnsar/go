package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	files, err := ioutil.ReadDir("../")
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}
