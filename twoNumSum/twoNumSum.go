package main

import "fmt"

//This is the optimal solution with time and space complexity of O(n)

//TwoNumberSum is the function to print the the numbers that sum up to give the target sum
func TwoNumberSum(array []int, target int) []int {
	numsMap := map[int]bool{}
	for _, num := range array {
		match := target - num
		if _, found := numsMap[match]; found {
			return []int{match, num}
		}
		numsMap[num] = true
	}
	return []int{}
}

func main() {
	arr := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	fmt.Print(TwoNumberSum(arr, target))
}
