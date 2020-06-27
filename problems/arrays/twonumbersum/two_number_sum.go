package twonumbersum

import "sort"

// // TwoNumberSum has O(n^2) time and O(1) space [Version 1]
// func TwoNumberSum(array []int, target int) []int {
// 	for i := 0; i < len(array)-1; i++ {
// 		firstNum := array[i]
// 		for j := i + 1; j < len(array); j++ {
// 			secondNum := array[j]
// 			if firstNum+secondNum == target {
// 				return []int{firstNum, secondNum}
// 			}
// 		}
// 	}
// 	return []int{}
// }

// // TwoNumberSum has O(n) time and O(n) space [Version 2]
// func TwoNumberSum(array []int, target int) []int {
// 	nums := map[int]bool{}
// 	for _, num := range array {
// 		potentialMatch := target - num
// 		if _, found := nums[potentialMatch]; found {
// 			return []int{potentialMatch, num}
// 		}
// 		nums[num] = true
// 	}
// 	return []int{}
// }

// TwoNumberSum has O(n) time and O(n) space [Version 3]
func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		currentSum := array[left] + array[right]
		if currentSum == target {
			return []int{array[left], array[right]}
		} else if currentSum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}
