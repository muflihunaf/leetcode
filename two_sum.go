package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	// map to store the number and its index
	mapTwoSum := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		// complement is the number that when added to the current number, equals the target
		complement := target - nums[i]
		// if the complement is in the map, return the index of the complement and the current index
		if index, ok := mapTwoSum[complement]; ok {
			return []int{index, i}
		}
		// if the complement is not in the map, add the current number and its index to the map
		mapTwoSum[nums[i]] = i
	}
	// if no solution is found, return nil
	return nil
}
