package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	// map to store the last index of each character
	mapData := make(map[byte]int)
	// left pointer
	left := 0
	// result
	res := 0

	// iterate over the string
	for i := 0; i < len(s); i++ {
		// if the character is already in the map, move the left pointer to the next character
		if _, ok := mapData[s[i]]; ok {
			left = max(left, mapData[s[i]]+1)
		}
		// update the last index of the character
		mapData[s[i]] = i
		// update the result
		res = max(res, i-left+1)
	}
	return res
}
