package main

import "fmt"

type nameAge struct {
	name string
	age  int
}

func main() {
	var nameAgeSlice []nameAge
	nameAges := map[string]int{
		"Michael":  34,
		"Franklin": 28,
		"Trevor":   43,
		"Kalam":    14,
	}

	for key, val := range nameAges {
		nameAgeSlice = append(nameAgeSlice, nameAge{key, val})
	}

	fmt.Println(nameAgeSlice)
}
