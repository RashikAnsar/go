package main

import (
	"fmt"
	"strings"
)

func main() {
	helloWorld := "Hello, World"
	helloUniverse := strings.Replace(helloWorld, "World", "Universe", 1)
	fmt.Println(helloUniverse)

	// arg 2 replaces first 2 occurances
	helloWorld = "Hello, World. How are you World?, I am good, thanks World"
	helloUniverse = strings.Replace(helloWorld, "World", "Universe", 2)
	fmt.Println(helloUniverse)

	// -1 replace all the occurances
	helloUniverse = strings.Replace(helloWorld, "World", "Universe", -1)
	fmt.Println(helloUniverse)
}
