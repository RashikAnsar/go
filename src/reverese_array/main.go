package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 3, 5, 9, 7, 6, 2, 10, 8, 4}
	toBeSorted := sort.IntSlice(numbers)
	// // To sort in increasing order
	// toBeSorted.Sort()
	// fmt.Println(toBeSorted)

	// To sort in decreasing order or reverse order
	sort.Sort(sort.Reverse(toBeSorted))
	fmt.Println(toBeSorted)

}
