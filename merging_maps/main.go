package main

import "fmt"

func main() {
	nameAges1 := map[string]int{
		"Michael":  34,
		"Franklin": 28,
		"Trevor":   43,
		"Kalam":    14,
	}
	fmt.Println(nameAges1)

	nameAges2 := map[string]int{
		"Benny":  30,
		"Amanda": 33,
		"Lester": 46,
	}

	for key, value := range nameAges2 {
		nameAges1[key] = value
	}

	fmt.Println(nameAges1)
}
