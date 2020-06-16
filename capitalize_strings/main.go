package main

import (
	"fmt"
	"strings"
)

func main() {
	helloWorld := "hello, how are you?"
	fmt.Println(helloWorld)

	helloWorldTitle := strings.Title(helloWorld)
	fmt.Println(helloWorldTitle)

	helloWorldUpper := strings.ToUpper(helloWorld)
	fmt.Println(helloWorldUpper)

	helloWorldLower := strings.ToLower(helloWorldUpper)
	fmt.Println(helloWorldLower)

}
