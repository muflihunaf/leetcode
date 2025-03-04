package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	nums2 := []int{0, 1, 3}
	fmt.Println(missingNumber(nums))
	fmt.Println(missingNumber(nums2))
}

func missingNumber(nums []int) int {
	// sum of array
	arraySum := 0
	for _, num := range nums {
		arraySum += num
	}
	// sum of total
	totalSum := len(nums) * (len(nums) + 1) / 2
	return totalSum - arraySum
}
