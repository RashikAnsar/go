package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []string{"Hyderabad", "New Delhi", "Chennai", "Bangalore", "Mumbai", "Gurgram", "Noida", "Kolkata"}

	for i, v := range str {
		if v == "Mumbai" {
			fmt.Println("Index of ", v, " is ", i)
		}
	}

	// Sorting and searching problems
	// sort.StringSlice
	// sort.IntSlice
	sortedList := sort.StringSlice(str)
	sortedList.Sort()
	fmt.Println(sortedList)

	idx := sortedList.Search("Mumbai")
	fmt.Println(idx)
}
