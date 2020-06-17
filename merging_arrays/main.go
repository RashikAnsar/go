package main

import "fmt"

func main() {
	arr1 := []int{3, 4}
	arr2 := []int{1, 2}

	result := append(arr1, arr2...)
	fmt.Println(result)
}
