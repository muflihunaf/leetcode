package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 10}
	nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(findDuplicate(nums))
	fmt.Println(findDuplicate(nums2))
}

func findDuplicate(nums []int) bool {
	mapDuplicate := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if mapDuplicate[nums[i]] {
			return true
		}
		mapDuplicate[nums[i]] = true
	}
	return false
}
