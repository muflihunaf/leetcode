package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(12321))
}

func isPalindrome(s int) bool {
	return s == reverse(s)
}

func reverse(s int) int {
	reversed := 0
	for s > 0 {
		reversed = reversed*10 + s%10
		s /= 10
	}
	return reversed
}
