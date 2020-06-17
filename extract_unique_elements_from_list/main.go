package main

import "fmt"

func main() {
	intSlice := []int{1, 5, 5, 5, 5, 7, 8, 6, 6, 6}
	uniqueIntSlice := unique(intSlice)
	fmt.Println(uniqueIntSlice)
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	uniqueElements := []int{}

	for _, entry := range intSlice {
		if _, val := keys[entry]; !val {
			keys[entry] = true
			uniqueElements = append(uniqueElements, entry)
		}
	}
	return uniqueElements
}
