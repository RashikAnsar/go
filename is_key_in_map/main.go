package main

import "fmt"

func main() {
	nameAges := map[string]int{
		"Michael":  34,
		"Franklin": 28,
		"Trevor":   43,
		"Kalam":    14,
	}

	// fmt.Println(nameAges["Franklin"])
	// fmt.Println(nameAges["Amanda"])

	if value, exists := nameAges["Franklin"]; exists {
		fmt.Println("Franklin is in map and value is: ", value)
	} else {
		fmt.Println("ERROR:: Given key doesn't exist in map")
	}
}
